package services

import (
    "fmt"
    "math/rand"
    "os"
    "time"
    "vch/internal/models"
)

func GenerateOrderCode() string {
    now := time.Now().Format("20060102")
    x := rand.Intn(1000000)
    return fmt.Sprintf("VCH%s-%06d", now, x)
}

func ComputeTotals(items []models.Item, sur []models.Surcharge) (subtotal, total float64) {
    for _, it := range items { subtotal += float64(it.Qty) * it.Price }
    total = subtotal
    for _, s := range sur { total += s.Amount }
    return
}

// Tracking base, default to https://vuachuyentien.com/tracking/
func TrackingURL(order string) string {
    base := os.Getenv("TRACKING_BASE")
    if base == "" { base = "https://vuachuyentien.com/tracking/" }
    if base[len(base)-1] != '/' { base += "/" }
    return base + order
}
