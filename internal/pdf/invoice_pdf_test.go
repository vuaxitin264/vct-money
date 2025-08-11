package pdf

import "testing"

func TestBuildInvoicePDF(t *testing.T) {
    payload := map[string]string{
        "qr": "https://vuachuyentien.com/tracking/VCH20250101-000001",
        "customer": "Test User\nPhone: 000\nAddress: X",
        "items": "- A x1 = $10",
        "total": "$10.00",
    }
    bin, err := BuildInvoicePDF("VCH20250101-000001", payload)
    if err != nil { t.Fatalf("BuildInvoicePDF error: %v", err) }
    if len(bin) < 500 { t.Fatalf("pdf too small: %d bytes", len(bin)) }
}
