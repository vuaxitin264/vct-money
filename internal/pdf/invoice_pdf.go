package pdf

import (
    "fmt"
    "image/png"
    "bytes"
    "github.com/jung-kurt/gofpdf"
    barcode "github.com/boombuler/barcode"
    code128 "github.com/boombuler/barcode/code128"
    qrcode "github.com/boombuler/barcode/qr"
)

// 4x6 inch label
func BuildInvoicePDF(orderCode string, payload map[string]string) ([]byte, error) {
    pdf := gofpdf.New("P", "mm", "Custom", "")
    pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 101.6, Ht: 152.4})
    pdf.SetFont("Arial", "", 11)

    pdf.CellFormat(0, 6, "Vua Chuyen Hang", "", 1, "C", false, 0, "")
    pdf.SetFont("Arial", "", 9)
    pdf.CellFormat(0, 5, fmt.Sprintf("Order: %s", orderCode), "", 1, "C", false, 0, "")

    // QR
    qr, _ := qrcode.Encode(payload["qr"], qrcode.M, qrcode.Auto)
    qr, _ = barcode.Scale(qr, 120, 120)
    var qb bytes.Buffer
    _ = png.Encode(&qb, qr)
    pdf.ImageOptionsReader(&qb, 5, 20, 30, 30, false, gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")

    // Barcode
    bc, _ := code128.Encode(orderCode)
    bc, _ = barcode.Scale(bc, 200, 40)
    var bb bytes.Buffer
    _ = png.Encode(&bb, bc)
    pdf.ImageOptionsReader(&bb, 5, 55, 80, 20, false, gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")

    // Customer + totals
    pdf.SetXY(40, 20)
    pdf.MultiCell(65, 5, payload["customer"], "", "L", false)

    pdf.SetXY(5, 80)
    pdf.MultiCell(95, 5, payload["items"], "", "L", false)

    pdf.SetXY(5, 120)
    pdf.SetFont("Arial", "B", 12)
    pdf.CellFormat(95, 8, fmt.Sprintf("TOTAL: %s", payload["total"]), "1", 0, "C", false, 0, "")

    var out bytes.Buffer
    if err := pdf.Output(&out); err != nil { return nil, err }
    return out.Bytes(), nil
}
