package cmd

import (
	"testing"
)

func Test_size2string(t *testing.T) {
	sz0 := 1 << 1        // 2B
	sz1 := 2 * (1 << 10) // 2K
	sz2 := 2 * (1 << 20) // 2M
	sz3 := 2 * (1 << 30) // 2G
	sz4 := 2 * (1 << 40) // 2T
	szStr := ""

	szStr = size2string(float64(sz0))
	if szStr != "2.00B" {
		t.Error("sz0 should be 2.00B")
	}

	szStr = size2string(float64(sz1))
	if szStr != "2.00K" {
		t.Error("sz0 should be 2.00K")
	}

	szStr = size2string(float64(sz2))
	if szStr != "2.00M" {
		t.Error("sz0 should be 2.00M")
	}

	szStr = size2string(float64(sz3))
	if szStr != "2.00G" {
		t.Error("sz0 should be 2.00G")
	}

	szStr = size2string(float64(sz4))
	if szStr != "2.00T" {
		t.Error("sz0 should be 2.00T")
	}
}
