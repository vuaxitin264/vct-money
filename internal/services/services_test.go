package services

import (
    "os"
    "regexp"
    "testing"
    "vch/internal/models"
)

func TestGenerateOrderCodeFormat(t *testing.T) {
    got := GenerateOrderCode()
    re := regexp.MustCompile("^VCH\\d{8}-\\d{6}$")
    if !re.MatchString(got) { t.Fatalf("order code format wrong: %s", got) }
}

func TestComputeTotals(t *testing.T) {
    items := []models.Item{{Name: "A", Qty: 2, Price: 5.0}, {Name: "B", Qty: 1, Price: 10.0}}
    sur := []models.Surcharge{{Label: "Fuel", Amount: 3.0}, {Label: "Handling", Amount: 2.0}}
    subtotal, total := ComputeTotals(items, sur)
    if subtotal != 20.0 { t.Fatalf("subtotal = %v, want 20.0", subtotal) }
    if total != 25.0 { t.Fatalf("total = %v, want 25.0", total) }
}

func TestTrackingURL(t *testing.T) {
    os.Setenv("TRACKING_BASE", "https://vuachuyentien.com/tracking/")
    got := TrackingURL("VCH20250101-000001")
    want := "https://vuachuyentien.com/tracking/VCH20250101-000001"
    if got != want { t.Fatalf("got %s, want %s", got, want) }
}
