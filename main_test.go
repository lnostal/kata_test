package main

import (
	"testing"
)

func TestConvToRoman(t *testing.T) {
	got, _ := convertArabicToRoman(89)
	want := "LXXXIX"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestConvToArabic(t *testing.T) {
	got, _ := convertRomanToArabic("IV")
	want := 4

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
