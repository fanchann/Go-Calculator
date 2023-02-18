package main

import (
	"Go-Calc/repository"
	"fmt"
)

func main() {
	calc := repository.Calculator{}
	calculator := repository.NewCalculator(&calc)

	result := calculator.Add(12, 23)
	fmt.Println(result)

	result = calculator.Multiply(12, 23)
	fmt.Println(result)

	result = int(calculator.Divide(20, 5))
	fmt.Println(result)

	result = int(calculator.Less(4, -5))
	fmt.Println(result)
}
