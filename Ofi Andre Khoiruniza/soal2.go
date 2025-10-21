package main

import "fmt"

func main() {
	var bmi, tinggiBadan float64
	fmt.Print("Input, BMI dan Tinggi Badan : ")
	fmt.Scan(&bmi)
	fmt.Scanln(&tinggiBadan)
	fmt.Printf("Output, Berat Badan : %.2f", bmi*(tinggiBadan*tinggiBadan))
}