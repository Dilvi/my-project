package main

import "fmt"

func main() {
	fmt.Println("=== Конвертер валют ===")

	from := inputCurrencyFrom()
	amount := inputAmount()
	to := inputCurrencyTo(from)

	result := currencyConverter(from, to, amount)

	fmt.Printf("Результат: %.2f\n", result)
}


func inputCurrencyFrom() int {
	var from int
	for {
		fmt.Println("Выберите исходную валюту:")
		fmt.Println("1 - EUR")
		fmt.Println("2 - USD")
		fmt.Println("3 - RUB")
		fmt.Print("Ввод: ")
		fmt.Scan(&from)
		if from >= 1 && from <= 3 {
			break
		}
		fmt.Println("Ошибка: введите 1, 2 или 3.")
	}
	return from
}


func inputAmount() float64 {
	var amount float64
	for {
		fmt.Println("Введите сумму:")
		fmt.Print("Ввод: ")
		fmt.Scan(&amount)
		if amount > 0 {
			break
		}
		fmt.Println("Сумма должна быть положительной.")
	}
	return amount
}


func inputCurrencyTo(from int) int {
	var choice int
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
		fmt.Print("Ввод: ")
		fmt.Scan(&choice)
		if choice == 1 || choice == 2 {
			break
		}
		fmt.Println("Ошибка: выберите 1 или 2.")
	}

	
	switch from {
	case 1: 
		if choice == 1 {
			return 2 
		}
		return 3 
	case 2: 
		if choice == 1 {
			return 1 
		}
		return 3 
	case 3: 
		if choice == 1 {
			return 1 
		}
		return 2 
	}
	return 0
}


func currencyConverter(from int, to int, amount float64) float64 {
	const USD_TO_EUR = 0.85
	const USD_TO_RUB = 78.47

	
	var usd float64
	switch from {
	case 1: 
		usd = amount / USD_TO_EUR
	case 2: 
		usd = amount
	case 3: 
		usd = amount / USD_TO_RUB
	}


	var result float64
	switch to {
	case 1: 
		result = usd * USD_TO_EUR
	case 2: 
		result = usd
	case 3: 
		result = usd * USD_TO_RUB
	}

	return result
}
