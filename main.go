package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Калькулятор операций")
	operation := operationInput()
	notFormattedValues := valuesInput()
	formattedValues := formatInput(notFormattedValues)
	for len(formattedValues) == 0 {
		fmt.Println("Не удалось прочитать ни одного числа. Повторите ввод.")
		notFormattedValues = valuesInput()
		formattedValues = formatInput(notFormattedValues)
	}
	result := calculate(formattedValues, operation)
	fmt.Println("Результат: ", result)
}

func operationInput() string {
	var operation string
	for {
		fmt.Println("Введите желаемую операцию:(AVG - среднее, SUM - сумма, MED - медиана)")
		fmt.Scanln(&operation)
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			break
		} else {
			fmt.Println("Операция не распознана, повторите ввод")
		}
	}
	return operation
}

func valuesInput() string {
	var notFormattedValues string
	fmt.Println("Введите числа для проведения операции через запятую: (Например 1, 2, 3...)")
	fmt.Scanln(&notFormattedValues)
	return notFormattedValues
}

func formatInput(notFormattedValues string) []float64 {
	formattedValues := []float64{}
	stringValues := strings.Split(notFormattedValues, ",")
	for index := range stringValues {
		cleanStringValue := strings.TrimSpace(stringValues[index])
		if cleanStringValue == "" {
			continue
		}
		intValue, err := strconv.ParseFloat(cleanStringValue, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		formattedValues = append(formattedValues, intValue)
	}
	return formattedValues
}

func calculate(formattedvalues []float64, operation string) float64 {
	var result float64
	switch {
	case operation == "AVG":
		var sum float64
		for index := range formattedvalues {
			sum += formattedvalues[index]
		}
		result = sum / float64(len(formattedvalues))
	case operation == "SUM":
		for index := range formattedvalues {
			result += formattedvalues[index]
		}
	case operation == "MED":
		sort.Float64s(formattedvalues)
		n := len(formattedvalues)
		if n % 2 == 0 {
			result = (formattedvalues[n/2-1] + formattedvalues[n/2])/2
		} else {
			result = formattedvalues[n/2]
		}
	}
	return result
}
