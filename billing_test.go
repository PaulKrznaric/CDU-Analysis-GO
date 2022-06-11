package main

import (
	"testing"
)

func TestFormatBillingNumber(t *testing.T) {
	got := FormatBillingNumber("J00123")
	want := 123
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
