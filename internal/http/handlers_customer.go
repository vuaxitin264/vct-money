package http

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/gofiber/fiber/v2"
    "vch/internal/models"
    "vch/internal/repositories"
)

type CustomerHandler struct{ Repo *repositories.CustomerRepo }

func (h *CustomerHandler) Dashboard(c *fiber.Ctx) error {
    return c.Render("dashboard", fiber.Map{"title":"Dashboard"})
}

func (h *CustomerHandler) List(c *fiber.Ctx) error {
    q := c.Query("q")
    list, err := h.Repo.List(c.Context(), q, 50, 0)
    if err != nil { return c.Status(500).SendString(err.Error()) }
    return c.Render("customers", fiber.Map{"customers": list, "q": q, "title":"Customers"})
}

func (h *CustomerHandler) Create(c *fiber.Ctx) error {
    var p struct {
        FullName string `form:"FullName"`
        Phone string `form:"Phone"`
        Email string `form:"Email"`
        Address string `form:"Address"`
        DateOfBirth string `form:"DateOfBirth"`
        DriverLicense string `form:"DriverLicense"`
    }
    if err := c.BodyParser(&p); err != nil { return c.Status(400).SendString("bad request") }

    rand.Seed(time.Now().UnixNano())
    code := "VCH" + fmt.Sprintf("%06d", rand.Intn(1000000))

    cst := &models.Customer{ Code: code, FullName: p.FullName, Phone: p.Phone, Email: p.Email, Address: p.Address, DateOfBirth: p.DateOfBirth, DriverLicense: p.DriverLicense }
    if err := h.Repo.Create(c.Context(), cst); err != nil { return c.Status(500).SendString(err.Error()) }
    return c.Redirect("/customers")
}
