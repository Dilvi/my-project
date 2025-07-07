package main

import "fmt"

func main() {
	fmt.Println("Конвертер валюты")
	currencyInput, currencyOutput := inputCurrency()
	currencyValue := inputAmount()
	result := currencyConverter(currencyInput, currencyValue, currencyOutput)
	fmt.Printf("Результат: %.2f", result)
}

func inputAmount() float64 {
	var currencyValue float64
	for {
		fmt.Println("Введите конвертируемую сумму")
		fmt.Scan(&currencyValue)
		if currencyValue <= 0 {
			fmt.Println("Ошибка при вводе, повторите ввод")
		} else {
			break
		}
	}
	return currencyValue
}

func inputCurrency() (int, int) {
	var currencyInput int
	var currencyOutput int
	for {
		fmt.Println("Введите конвертируемую валюту")
		fmt.Println("1 - EUR")
		fmt.Println("2 - USD")
		fmt.Println("3 - RUB")
		fmt.Scan(&currencyInput)
		if currencyInput != 1 && currencyInput != 2 && currencyInput != 3 {
			fmt.Println("Ошибка при вводе, повторите ввод")
		} else {
			break
		}
	}
	for {
		fmt.Println("В какую валюту вы хотите конвертировать?")
		switch {
		case currencyInput == 1:
			fmt.Println("1 - USD")
			fmt.Println("2 - RUB")
		case currencyInput == 2:
			fmt.Println("1 - EUR")
			fmt.Println("2 - RUB")
		case currencyInput == 3:
			fmt.Println("1 - EUR")
			fmt.Println("2 - USD")
		}
		fmt.Scan(&currencyOutput)
		if currencyOutput != 1 && currencyOutput != 2 {
			fmt.Println("Ошибка при вводе, повторите ввод")
		} else {
			break
		}
	}

	// Преобразуем выбор пользователя в реальные валюты
	switch currencyInput {
	case 1: // EUR
		if currencyOutput == 1 {
			currencyOutput = 2 // USD
		} else {
			currencyOutput = 3 // RUB
		}
	case 2: // USD
		if currencyOutput == 1 {
			currencyOutput = 1 // EUR
		} else {
			currencyOutput = 3 // RUB
		}
	case 3: // RUB
		if currencyOutput == 1 {
			currencyOutput = 1 // EUR
		} else {
			currencyOutput = 2 // USD
		}
	}
	return currencyInput, currencyOutput
}

func currencyConverter(currencyInput int, currencyValue float64, currencyOutput int) float64 {
	var result float64
	const USD_TO_EUR = 0.85
	const USD_TO_RUB = 78.47

	switch {
	case currencyInput == 1 && currencyOutput == 2:
		// EUR -> USD
		result = currencyValue * (1 / USD_TO_EUR)
	case currencyInput == 1 && currencyOutput == 3:
		// EUR -> RUB = EUR -> USD -> RUB
		result = currencyValue * (1 / USD_TO_EUR) * USD_TO_RUB
	case currencyInput == 2 && currencyOutput == 1:
		// USD -> EUR
		result = currencyValue * USD_TO_EUR
	case currencyInput == 2 && currencyOutput == 3:
		// USD -> RUB
		result = currencyValue * USD_TO_RUB
	case currencyInput == 3 && currencyOutput == 1:
		// RUB -> EUR = RUB -> USD -> EUR
		result = currencyValue * (1 / USD_TO_RUB) * USD_TO_EUR
	case currencyInput == 3 && currencyOutput == 2:
		// RUB -> USD
		result = currencyValue * (1 / USD_TO_RUB)
	}
	return result
}
