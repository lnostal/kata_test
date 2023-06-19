package main

const alert = "Ошибка: "

const (
	notMathsExpression         = alert + "Строка не является математической операцией"
	notInTheRange              = alert + "%d ∉ (0,10)"
	moreThenOneOperation       = alert + "В выражении больше двух операндов"
	notNumber                  = alert + "%s не является допустимым числом"
	unknown                    = alert + "¯\\_(ツ)_/¯"
	onlyPositiveNumbersInRoman = alert + "%d < 1"
	differentSystems           = alert + "Используются одновременно разные системы счисления"
)
