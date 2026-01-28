package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("__ Calculate your IMT __")

	for {
		userKg, userHeight, err := getUserInput()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()

		if !isRepeatCalculation {
			break
		}
	}

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Your IMT is: %.0f", IMT)
	fmt.Println(result)

	switch {
	case IMT < 16:
		fmt.Println("You are severely underweight")
	case IMT < 18.5:
		fmt.Println("You are underweight")
	case IMT < 25:
		fmt.Println("You are normal weight")
	case IMT < 30:
		fmt.Println("You are overweight")
	default:
		fmt.Println("You are obese")
	}
}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("weight and height must be positive numbers")
	}

	IMTPower := 2.0

	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64, error) {
	var userHeight float64
	var userKg float64
	fmt.Println("Enter your height in centimeters:")
	if _, err := fmt.Scan(&userHeight); err != nil {
		return 0, 0, errors.New("Invalid input: height must be a number")
	}
	fmt.Println("Enter your weight in kilograms:")
	if _, err := fmt.Scan(&userKg); err != nil {
		return 0, 0, errors.New("Invalid input: weight must be a number")
	}
	return userKg, userHeight, nil
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Println("Do you want to calculate your IMT again? (y/n):")
	fmt.Scan(&userChoice)
	return strings.ToLower(userChoice) == "y"
}
