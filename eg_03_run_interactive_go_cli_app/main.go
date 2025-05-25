package main

import (
	"fmt"
	"os"
)

func main() {
	var chosenSystem string
	fmt.Println("(1) Metric (m, kg) or (2) Non-Metric (ft, pounds)?")
	fmt.Print("Please choose: ")
	fmt.Scanln(&chosenSystem)

	if chosenSystem != "1" && chosenSystem != "2" {
		fmt.Println("You have to choose either metric or non-metric. Shutting down...")
		os.Exit(1)
	}

	heightUnit := "meters"
	weightUnit := "kilograms"
	if chosenSystem == "2" {
		heightUnit = "feet"
		weightUnit = "pounds"
	}

	var userHeight, userWeight float64
	fmt.Printf("Please enter your height in %s\n", heightUnit)
	fmt.Print("Your height: ")
	_, err := fmt.Scanln(&userHeight)
	if err != nil {
		fmt.Println("Invalid height input:", err)
		os.Exit(1)
	}

	fmt.Printf("Please enter your weight in %s\n", weightUnit)
	fmt.Print("Your weight: ")
	_, err = fmt.Scanln(&userWeight)
	if err != nil {
		fmt.Println("Invalid weight input:", err)
		os.Exit(1)
	}

	adjHeight := userHeight
	adjWeight := userWeight
	if chosenSystem == "2" {
		adjHeight = userHeight / 3.28084
		adjWeight = userWeight / 2.20462
	}

	bmi := adjWeight / (adjHeight * adjHeight)
	fmt.Printf("Your body-mass-index: %.10f\n", bmi) // show more decimals for clarity
}
