package main

import "fmt"

func main() {
	const USD = 78.47
	const EUR = 92.28
	EUR_value := 1.0
	EUR_to_USD := USD / EUR
	EUR_in_RUB := EUR_value * EUR_to_USD * USD
	fmt.Print(EUR_in_RUB)
}