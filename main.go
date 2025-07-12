package main

import (
	"fmt"
)

func main() {
	fmt.Println("Конвертер валюты")
	currencyInput, currencyOutput := inputCurrency()
	currencyValue := inputAmount()
	result := currencyConverter(currencyInput, currencyValue, currencyOutput)
	fmt.Printf("Результат: %.2f\n", result)
}

// Карта курсов: от какой валюты → в какую валюту → курс
var exchangeRates = map[string]map[string]float64{
	"EUR": {
		"USD": 1 / 0.85,               // EUR → USD
		"RUB": 78.47 / 0.85,           // EUR → RUB
	},
	"USD": {
		"EUR": 0.85,                   // USD → EUR
		"RUB": 78.47,                  // USD → RUB
	},
	"RUB": {
		"EUR": 0.85 / 78.47,           // RUB → EUR
		"USD": 1 / 78.47,              // RUB → USD
	},
}

func inputAmount() float64 {
	var currencyValue float64
	for {
		fmt.Println("Введите конвертируемую сумму:")
		fmt.Scan(&currencyValue)
		if currencyValue <= 0 {
			fmt.Println("Ошибка при вводе, повторите ввод.")
		} else {
			break
		}
	}
	return currencyValue
}

func inputCurrency() (string, string) {
	var currencies = map[int]string{
		1: "EUR",
		2: "USD",
		3: "RUB",
	}
	var currencyInput, currencyOutput int

	for {
		fmt.Println("Выберите валюту, из которой конвертировать:")
		fmt.Println("1 - EUR\n2 - USD\n3 - RUB")
		fmt.Scan(&currencyInput)
		if _, ok := currencies[currencyInput]; !ok {
			fmt.Println("Ошибка при вводе, повторите ввод.")
		} else {
			break
		}
	}

	for {
		fmt.Println("Выберите валюту, в которую конвертировать:")
		targets := []int{}
		for i := 1; i <= 3; i++ {
			if i != currencyInput {
				fmt.Printf("%d - %s\n", i, currencies[i])
				targets = append(targets, i)
			}
		}
		fmt.Scan(&currencyOutput)
		if currencyOutput == currencyInput || currencyOutput < 1 || currencyOutput > 3 {
			fmt.Println("Ошибка при вводе, повторите ввод.")
		} else {
			break
		}
	}

	return currencies[currencyInput], currencies[currencyOutput]
}

func currencyConverter(from string, amount float64, to string) float64 {
	rate := exchangeRates[from][to]
	return amount * rate
}
