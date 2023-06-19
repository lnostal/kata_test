package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func preCheckExpression(expression string) error {
	splited := strings.Split(expression, "")

	if len(splited) < 3 {
		return errors.New(notMathsExpression)
	}

	countEmpties := 0

	for _, s := range splited {
		if s == " " {
			countEmpties++
		}
	}

	if len(splited)-countEmpties < 3 {
		return errors.New(notMathsExpression)
	}

	return nil
}

func readExpression() string {

	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	//splited := strings.Split(expression, "")
	//
	//if len(splited) < 3 {
	//	return "", errors.New(notMathsExpression)
	//}
	//
	//countEmpties := 0
	//
	//for _, s := range splited {
	//	if s == " " {
	//		countEmpties++
	//	}
	//}
	//
	//if len(splited)-countEmpties < 3 {
	//	return "", errors.New(notMathsExpression)
	//}

	return expression
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

func analyzeExpression(s string, op operation) (a int, b int, o operation, rom bool, err error) {

	noErrors := preCheckExpression(s)
	if noErrors != nil {
		return 0, 0, op, false, errors.New(noErrors.Error())
	}

	numbers := strings.Split(s, string(op))

	for i := range numbers {
		numbers[i] = strings.TrimSpace(numbers[i])
	}

	// проверяем, что операндов не больше двух
	if len(numbers) > 2 {
		return 0, 0, op, false, errors.New(notMathsExpression)
	}

	// проверяем, не спряталась ли еще одна операция в выражении
	_, err = getOperator(numbers[1])

	if err == nil {
		return 0, 0, op, false, errors.New(moreThenOneOperation)
	}

	var numInts []int
	romanCounts := 0

	// проверяем, что все операнды - числа и не больше 10
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {

			// проверяем, это Roman или нет
			roman, n := isRoman(numStr)
			if roman {
				numInts = append(numInts, n)
				romanCounts++
				continue
			}

			return 0, 0, op, false, errors.New(fmt.Sprintf(notNumber, numStr))
		}

		if num > 10 || num < 1 {
			return 0, 0, op, false, errors.New(fmt.Sprintf(notInTheRange, num))
		}

		numInts = append(numInts, num)
	}

	if romanCounts == 1 {
		return 0, 0, op, false, errors.New(differentSystems)
	}

	return numInts[0], numInts[1], op, romanCounts == 2, nil
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

func isRoman(s string) (bool, int) {
	num, err := convertRomanToArabic(s)

	if err != nil {
		return false, 0
	}

	return true, num
}
