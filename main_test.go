package main

import (
	"fmt"
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

func Test_GetTwoArabicNumerals(t *testing.T) {
	got, _ := calc("1 + 2")
	want := "3"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test_GetTwoRomanNumerals(t *testing.T) {
	got, _ := calc("VI / III")
	want := "II"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test_GetTwoRomanNumeralsButAnswerBelowZero(t *testing.T) {
	_, err := calc("I - II")
	want := fmt.Sprintf(onlyPositiveNumbersInRoman, -1)

	if err.Error() != want {
		t.Errorf("got %q, wanted %q", err.Error(), want)
	}
}

func Test_GetDifferentSystems(t *testing.T) {
	_, err := calc("I + 1")
	want := fmt.Sprintf(differentSystems)

	if err.Error() != want {
		t.Errorf("got %q, wanted %q", err.Error(), want)
	}
}

func Test_GetNotMathsExpression(t *testing.T) {
	_, err := calc("1")
	want := fmt.Sprintf(notMathsExpression)

	if err.Error() != want {
		t.Errorf("got %q, wanted %q", err.Error(), want)
	}
}

func Test_GetMoreThenOneOperations(t *testing.T) {
	_, err := calc("1 + 2 + 3")
	want := fmt.Sprintf(moreThenOneOperation)

	if err.Error() != want {
		t.Errorf("got %q, wanted %q", err.Error(), want)
	}
}
