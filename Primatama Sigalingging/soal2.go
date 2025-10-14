package main

import "fmt"

func main() {
	var bmi, tinggiBadan, beratBadan float64

	fmt.Scan(&bmi, &tinggiBadan)

	beratBadan = bmi * (tinggiBadan * tinggiBadan)

	fmt.Printf("%.0f\n", beratBadan)
}