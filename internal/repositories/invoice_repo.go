package repositories

import (
    "context"
    "database/sql"
    "encoding/json"
    "vch/internal/models"
)

type InvoiceRepo struct{ DB *sql.DB }

func (r *InvoiceRepo) Create(ctx context.Context, inv *models.Invoice) error {
    itemsJSON, _ := json.Marshal(inv.Items)
    surJSON, _ := json.Marshal(inv.Surcharges)
    return r.DB.QueryRowContext(ctx,
        `INSERT INTO invoices(order_code,customer_id,items,weight_lbs,subtotal,surcharges,total,currency,transfer_amount,transfer_fee,total_recipient,invoice_date,bank_info_vn,status,notes)
         VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING id`,
        inv.OrderCode, inv.CustomerID, itemsJSON, inv.WeightLbs, inv.Subtotal, surJSON, inv.Total, inv.Currency,
        inv.TransferAmount, inv.TransferFee, inv.TotalRecipient, inv.InvoiceDate, inv.BankInfoVN, inv.Status, inv.Notes,
    ).Scan(&inv.ID)
}

func (r *InvoiceRepo) GetByID(ctx context.Context, id int64) (*models.Invoice, error) {
    row := r.DB.QueryRowContext(ctx,
        `SELECT id,order_code,customer_id,items,weight_lbs,subtotal,surcharges,total,currency,transfer_amount,transfer_fee,total_recipient,invoice_date,bank_info_vn,status,notes,created_at,updated_at
         FROM invoices WHERE id=$1`, id)
    var inv models.Invoice
    var itemsJSON, surJSON []byte
    if err := row.Scan(&inv.ID,&inv.OrderCode,&inv.CustomerID,&itemsJSON,&inv.WeightLbs,&inv.Subtotal,&surJSON,&inv.Total,&inv.Currency,&inv.TransferAmount,&inv.TransferFee,&inv.TotalRecipient,&inv.InvoiceDate,&inv.BankInfoVN,&inv.Status,&inv.Notes,&inv.CreatedAt,&inv.UpdatedAt); err != nil { return nil, err }
    _ = json.Unmarshal(itemsJSON, &inv.Items)
    _ = json.Unmarshal(surJSON, &inv.Surcharges)
    return &inv, nil
}
