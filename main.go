package main

import "fmt"

func main() {
	fmt.Println("Конвертер валюты")

	from := inputCurrencyFrom()
	amount := inputAmount()
	to := inputCurrencyTo(from)

	result := currencyConverter(from, amount, to)
	fmt.Printf("Результат: %.2f\n", result)
}

func inputCurrencyFrom() int {
	var currency int
	for {
		fmt.Println("Выберите исходную валюту:")
		fmt.Println("1 - EUR")
		fmt.Println("2 - USD")
		fmt.Println("3 - RUB")
		fmt.Scan(&currency)
		if currency >= 1 && currency <= 3 {
			break
		}
		fmt.Println("Ошибка при вводе, повторите.")
	}
	return currency
}

func inputAmount() float64 {
	var amount float64
	for {
		fmt.Println("Введите сумму для конвертации:")
		fmt.Scan(&amount)
		if amount > 0 {
			break
		}
		fmt.Println("Сумма должна быть больше 0.")
	}
	return amount
}

func inputCurrencyTo(from int) int {
	var choice int
	var to int
	for {
		fmt.Println("Выберите целевую валюту:")
		switch from {
		case 1:
			fmt.Println("1 - USD")
			fmt.Println("2 - RUB")
		case 2:
			fmt.Println("1 - EUR")
			fmt.Println("2 - RUB")
		case 3:
			fmt.Println("1 - EUR")
			fmt.Println("2 - USD")
		}
		fmt.Scan(&choice)
		if choice == 1 || choice == 2 {
			break
		}
		fmt.Println("Ошибка при вводе, повторите.")
	}
	switch from {
	case 1:
		if choice == 1 {
			to = 2 // EUR -> USD
		} else {
			to = 3 // EUR -> RUB
		}
	case 2:
		if choice == 1 {
			to = 1 // USD -> EUR
		} else {
			to = 3 // USD -> RUB
		}
	case 3:
		if choice == 1 {
			to = 1 // RUB -> EUR
		} else {
			to = 2 // RUB -> USD
		}
	}
	return to
}

func currencyConverter(from int, amount float64, to int) float64 {
	const USD_TO_EUR = 0.85
	const USD_TO_RUB = 78.47

	var usd float64

	// Сначала переводим всё в USD
	switch from {
	case 1: // EUR -> USD
		usd = amount * (1 / USD_TO_EUR)
	case 2: // USD
		usd = amount
	case 3: // RUB -> USD
		usd = amount * (1 / USD_TO_RUB)
	}

	var result float64
	// Потом из USD в нужную валюту
	switch to {
	case 1: // EUR
		result = usd * USD_TO_EUR
	case 2: // USD
		result = usd
	case 3: // RUB
		result = usd * USD_TO_RUB
	}

	return result
}
