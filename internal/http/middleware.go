package http

import (
    "os"
    "github.com/gofiber/fiber/v2"
)

type Middleware struct{}

func (m *Middleware) Auth(c *fiber.Ctx) error {
    if c.Cookies("admin") == "1" {
        return c.Next()
    }
    return c.Redirect("/login")
}

func (m *Middleware) RequirePassword(pw string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        if c.Cookies("admin") == "1" { return c.Next() }
        if c.Method() == fiber.MethodPost {
            if c.FormValue("password") == pw {
                c.Cookie(&fiber.Cookie{Name: "admin", Value: "1", Path: "/", HTTPOnly: true, Secure: false})
                return c.Redirect("/dashboard")
            }
            return c.Render("login", fiber.Map{"error": "Wrong password", "title":"Login"})
        }
        return c.Render("login", fiber.Map{"title":"Login"})
    }
}

func AdminPassword() string { if v := os.Getenv("ADMIN_PASSWORD"); v != "" { return v }; return "changeme" }
