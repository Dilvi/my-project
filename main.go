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
	fmt.Scan(EURValue)
	return EURValue
}

// Пустая функция без реализации
func currencyConverter(a,b,c float64) float64 { // Принимает 3 значения: 1 число и 2 валюты - исходная и целевая
	var d float64 // Выходное число
	return d
}