package main

import (
	"fmt"
	"strconv"
)

func calc(text string) (string, error) {
	op, err := getOperator(text)

	if err != nil {
		return "", err
	}

	a, b, op, isRoman, err := analyzeExpression(text, op)

	if err != nil {
		return "", err
	}

	answer, _ := calculate(a, b, op)

	if isRoman {
		ansStr, err := convertArabicToRoman(answer)
		if err != nil {
			return ansStr, err
		}

		return ansStr, nil
	}

	return strconv.Itoa(answer), nil
}

func main() {

	fmt.Println("Введите выражение")
	text := readExpression()

	answer, err := calc(text)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Ответ: %s\n", answer)
}
