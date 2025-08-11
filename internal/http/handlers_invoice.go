package http

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/gofiber/fiber/v2"
    "vch/internal/models"
    "vch/internal/pdf"
    "vch/internal/repositories"
    "vch/internal/services"
)

type InvoiceHandler struct{ Repo *repositories.InvoiceRepo }

func (h *InvoiceHandler) List(c *fiber.Ctx) error {
    return c.Render("invoices", fiber.Map{"title":"Invoices"})
}

func (h *InvoiceHandler) Create(c *fiber.Ctx) error {
    var req struct {
        CustomerID int64              `json:"customer_id" form:"customer_id"`
        Items      []models.Item      `json:"items"`
        Surcharges []models.Surcharge `json:"surcharges"`
        WeightLbs  float64            `json:"weight_lbs" form:"weight_lbs"`
        Notes      string             `json:"notes" form:"notes"`
        Currency   string             `form:"currency"`
        TransferAmount float64        `form:"transfer_amount"`
        TransferFee    float64        `form:"transfer_fee"`
        TotalRecipient float64        `form:"total_recipient"`
        InvoiceDate    string         `form:"invoice_date"`
        BankInfoVN     string         `form:"bank_info_vn"`
    }
    if err := c.BodyParser(&req); err != nil { return c.Status(400).SendString("bad request") }

    order := services.GenerateOrderCode()
    subtotal, total := services.ComputeTotals(req.Items, req.Surcharges)

    inv := &models.Invoice{
        OrderCode:  order,
        CustomerID: req.CustomerID,
        Items:      req.Items,
        WeightLbs:  req.WeightLbs,
        Subtotal:   subtotal,
        Surcharges: req.Surcharges,
        Total:      total,
        Currency:   strings.ToUpper(req.Currency),
        TransferAmount: req.TransferAmount,
        TransferFee:    req.TransferFee,
        TotalRecipient: req.TotalRecipient,
        InvoiceDate:    req.InvoiceDate,
        BankInfoVN:     req.BankInfoVN,
        Status:     "NEW",
        Notes:      req.Notes,
    }
    if err := h.Repo.Create(c.Context(), inv); err != nil { return c.Status(500).SendString(err.Error()) }

    // Redirect to HTML label route for DIRECT PRINT
    return c.Redirect(fmt.Sprintf("/invoices/%d/label?print=1", inv.ID))
}

func (h *InvoiceHandler) PDF(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, _ := strconv.ParseInt(idStr, 10, 64)
    inv, err := h.Repo.GetByID(c.Context(), id)
    if err != nil { return c.Status(404).SendString("invoice not found") }

    // build payload
    payload := map[string]string{
        "qr": services.TrackingURL(inv.OrderCode),
        "customer": "Order for customer #"+fmt.Sprint(inv.CustomerID),
        "items": "(items omitted)",
        "total": fmt.Sprintf("$%.2f", inv.Total),
    }

    bin, err := pdf.BuildInvoicePDF(inv.OrderCode, payload)
    if err != nil { return c.Status(500).SendString(err.Error()) }
    c.Set("Content-Type", "application/pdf")
    c.Set("Content-Disposition", "inline; filename=label-"+inv.OrderCode+".pdf")
    return c.Send(bin)
}

// HTML label 4x6 for immediate print via window.print()
func (h *InvoiceHandler) Label4x6(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, _ := strconv.ParseInt(idStr, 10, 64)
    inv, err := h.Repo.GetByID(c.Context(), id)
    if err != nil { return c.Status(404).SendString("invoice not found") }

    // formatting depending on currency
    currency := inv.Currency
    format := func(v float64) string {
        if currency == "VND" || currency == "VNĐ" {
            return fmt.Sprintf("%.0f VND", v)
        }
        return fmt.Sprintf("$%.2f", v)
    }

    printNow := strings.ToLower(c.Query("print")) == "1"
    return c.Render("label4x6", fiber.Map{
        "title": "Print Label",
        "order": inv.OrderCode,
        "tracking": services.TrackingURL(inv.OrderCode),
        "amount": format(inv.TransferAmount),
        "fee": format(inv.TransferFee),
        "total_recipient": format(inv.TotalRecipient),
        "invoice_date": inv.InvoiceDate,
        "bank_info_vn": inv.BankInfoVN,
        "currency": currency,
        "printNow": printNow,
    })
}
