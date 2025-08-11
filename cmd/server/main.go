package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    html "github.com/gofiber/template/html/v2"
    "vch/internal/db"
    vhttp "vch/internal/http"
    "vch/internal/repositories"
)

func main(){
    engine := html.New("./web/templates", ".html")

    database, err := db.Connect()
    if err != nil { log.Fatal(err) }
    if err := db.Migrate(database); err != nil { log.Fatal(err) }

    custRepo := &repositories.CustomerRepo{DB: database}
    invRepo := &repositories.InvoiceRepo{DB: database}

    app := fiber.New(fiber.Config{Views: engine})

    hc := &vhttp.CustomerHandler{Repo: custRepo}
    hi := &vhttp.InvoiceHandler{Repo: invRepo}
    mw := &vhttp.Middleware{}

    vhttp.MountRoutes(app, hc, hi, mw)

    app.Static("/static", "./web/static")

    log.Println("listening on :8080")
    if err := app.Listen(":8080"); err != nil { log.Fatal(err) }
}
