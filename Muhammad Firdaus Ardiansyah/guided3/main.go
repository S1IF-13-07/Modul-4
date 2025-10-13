package main

import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Print("Masukkan berat badan (kg) : ")
	fmt.Scan(&beratBadan)
	fmt.Print("Masukkan tinggi badan (m) : ")
	fmt.Scan(&tinggiBadan)
	bmi = beratBadan / (tinggiBadan * tinggiBadan)
	fmt.Printf("%.2f", bmi)
}