package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.85     
	const USD_TO_RUB = 78.47   
	EURValue := userInput()
	EUR_to_RUB := EURValue * (USD_TO_RUB / USD_TO_EUR)
	fmt.Print(EUR_to_RUB)
}

func userInput() float64 {
	var EURValue float64
	fmt.Println("Введите количество Евро, которые вы хотите перевести в рубли")
	fmt.Scan(&EURValue)
	return EURValue
}


func currencyConverter(amount float64, from string, to string) float64 {
	var result float64
	return result
}
