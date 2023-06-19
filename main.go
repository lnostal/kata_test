package main

import "fmt"

func main() {

	fmt.Println("Введите выражение")
	text, err := readExpression()

	if err != nil {
		returnError(err)
		return
	}

	op, err := getOperator(text)

	if err != nil {
		returnError(err)
		return
	}

	a, b, op, isRoman, err := analyzeExpression(text, op)

	if err != nil {
		returnError(err)
		return
	}

	answer, _ := calculate(a, b, op)

	if isRoman {
		ansStr, err := convertArabicToRoman(answer)
		if err != nil {
			returnError(err)
			return
		}

		fmt.Printf("Ответ: %s\n", ansStr)
		return
	}

	fmt.Printf("Ответ: %d\n", answer)
}
