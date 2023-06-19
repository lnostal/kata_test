package main

import (
	"errors"
	"fmt"
	"strings"
)

type Numerals struct {
	arabic int
	roman  string
}

var SetOfNumerals = [11]Numerals{
	{arabic: 0, roman: "s"},
	{arabic: 1, roman: "I"},
	{arabic: 2, roman: "II"},
	{arabic: 3, roman: "III"},
	{arabic: 4, roman: "IV"},
	{arabic: 5, roman: "V"},
	{arabic: 6, roman: "VI"},
	{arabic: 7, roman: "VII"},
	{arabic: 8, roman: "VIII"},
	{arabic: 9, roman: "IX"},
	{arabic: 10, roman: "X"},
}

func convertRomanToArabic(number string) (int, error) {
	for _, r := range SetOfNumerals {
		if number == r.roman {
			return r.arabic, nil
		}
	}
	return 0, errors.New(unknown)
}

func convertArabicToRoman(number int) (string, error) {

	if number < 1 {
		return "", errors.New(fmt.Sprintf(onlyPositiveNumbersInRoman, number))
	}

	if number < 11 {
		for _, r := range SetOfNumerals {
			if number == r.arabic {
				return r.roman, nil
			}
		}
		return "", errors.New(unknown)
	}

	if number < 100 {

		if number/10 < 4 {
			str, err := addOrder(number)
			return str, err
		}

		if number/10 == 4 {
			mod := number % 10
			if mod < 9 {
				str, _ := convertArabicToRoman(mod)
				return "XL" + str, nil
			}

			return "XLIX", nil
		}

		if number/10 < 9 {
			str, _ := addOrder(number - 50)
			return "L" + str, nil
		} else if number/10 == 9 {
			str, _ := addOrder(number - 90)
			return "XC" + str, nil
		}

	}

	return "C", nil
}

func addOrder(number int) (string, error) {
	for i := 0; i < 4; i++ {
		if number/10 == i {

			mod := number % 10
			if mod < 9 {
				str, _ := convertArabicToRoman(mod)
				return strings.Repeat("X", i) + str, nil
			}

			return strings.Repeat("X", i) + "IX", nil
		}
	}

	return "", nil
}
