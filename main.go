package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.85     
	const USD_TO_RUB = 78.47   
	EURValue := 1.0
	EUR_to_RUB := EURValue * (USD_TO_RUB / USD_TO_EUR)
	fmt.Print(EUR_to_RUB)
}
