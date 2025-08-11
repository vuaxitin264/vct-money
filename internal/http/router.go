package http

import (
    "github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App, hc *CustomerHandler, hi *InvoiceHandler, mw *Middleware) {
    app.Get("/", func(c *fiber.Ctx) error { return c.Redirect("/dashboard") })

    // auth
    app.All("/login", mw.RequirePassword(AdminPassword()))

    // protected
    app.Get("/dashboard", mw.Auth, hc.Dashboard)

    // Customers
    app.Get("/customers", mw.Auth, hc.List)
    app.Post("/customers", mw.Auth, hc.Create)

    // Invoices
    app.Get("/invoices", mw.Auth, hi.List)
    app.Post("/invoices", mw.Auth, hi.Create)
    app.Get("/invoices/:id/pdf", mw.Auth, hi.PDF)
    app.Get("/invoices/:id/label", mw.Auth, hi.Label4x6) // HTML 4x6 for direct print
}
