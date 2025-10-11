package main

import "fmt"

func main() {
	var bmi, tinggi, berat float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scanln(&bmi)
	fmt.Print("Masukkan tinggi badan : ")
	fmt.Scanln(&tinggi)

	berat = bmi * tinggi * tinggi

	fmt.Printf("%.0f",berat)
}