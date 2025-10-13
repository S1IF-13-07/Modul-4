package main

import "fmt"

func main() {
	var bmi, tinggi, berat float64
	fmt.Print("Masukkan nilai BMI dan tinggi badan (m): ")
	fmt.Scan(&bmi, &tinggi)
	berat = bmi * (tinggi * tinggi)
	fmt.Printf("Berat badan: %.0f\n", berat)
}
