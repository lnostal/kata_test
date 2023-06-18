package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation string

const (
	addition       = "+"
	subtraction    = "-"
	division       = "/"
	multiplication = "*"
)

const alert = "Ошибка: "

type errs string

const (
	notMathsExpression   = alert + "Строка не является математической операцией"
	notInTheRange        = alert + "%d ∉ (1,10)"
	moreThenOneOperation = alert + "В выражении больше двух операндов"
	notNumber            = alert + "%s не является числом"
	unknown              = alert + "¯\\_(ツ)_/¯"
)

func readExpression() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	splited := strings.Split(expression, "")

	if len(splited) < 3 {
		return "", errors.New(notMathsExpression)
	}

	countEmpties := 0

	for _, s := range splited {
		if s == " " {
			countEmpties++
		}
	}

	if len(splited)-countEmpties != 3 {
		return "", errors.New(notMathsExpression)
	}

	return expression, nil
}

func getOperator(s string) (operation, error) {
	set := []operation{addition, subtraction, division, multiplication}

	for _, op := range set {
		res := strings.Contains(s, string(op))
		if res == true {
			return op, nil
		}
	}

	return "", errors.New(notMathsExpression)
}

func analyzeExpression(s string, op operation) (a int, b int, o operation, err error) {

	numbers := strings.Split(s, string(op))

	for i := range numbers {
		numbers[i] = strings.TrimSpace(numbers[i])
	}

	// проверяем, что операндов не больше двух
	if len(numbers) > 2 {
		return 0, 0, op, errors.New(notMathsExpression)
	}

	// проверяем, не спряталась ли еще одна операция в выражении
	_, err = getOperator(numbers[1])

	if err == nil {
		return 0, 0, op, errors.New(moreThenOneOperation)
	}

	var numInts []int
	// проверяем, что все операнды - числа и не больше 10
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, 0, op, errors.New(fmt.Sprintf(notNumber, numStr))
		}

		if num > 10 || num < 1 {
			return 0, 0, op, errors.New(fmt.Sprintf(notInTheRange, num))
		}

		numInts = append(numInts, num)
	}

	return numInts[0], numInts[1], op, nil
}

func returnError(e error) {
	fmt.Println(e.Error())
	return
}

func calculate(a, b int, op operation) (int, error) {

	switch op {
	case addition:
		return a + b, nil

	case subtraction:
		return a - b, nil
	case division:
		return a / b, nil
	case multiplication:
		return a * b, nil
	}

	return 0, errors.New(unknown)
}

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

	a, b, op, err := analyzeExpression(text, op)

	if err != nil {
		returnError(err)
		return
	}

	answer, _ := calculate(a, b, op)
	fmt.Printf("Ответ: %d\n", answer)

}
