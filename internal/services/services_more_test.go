package services

import (
    "os"
    "testing"
)

func TestTrackingURLAddsSlash(t *testing.T) {
    os.Setenv("TRACKING_BASE", "https://vuachuyentien.com/tracking")
    got := TrackingURL("ABC")
    want := "https://vuachuyentien.com/tracking/ABC"
    if got != want { t.Fatalf("got %s, want %s", got, want) }
}

func TestComputeTotalsEmpty(t *testing.T) {
    s, ttotal := ComputeTotals(nil, nil)
    if s != 0 || ttotal != 0 {
        t.Fatalf("empty totals should be 0,0 got %v,%v", s, ttotal)
    }
}
