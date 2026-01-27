package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Calculate your IMT __")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Your IMT is: %.0f", IMT)
	fmt.Println(result)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Println("Enter your height in santimeters:")
	fmt.Scan(&userHeight)
	fmt.Println("Enter your weight in kilograms:")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
