package main

import "fmt"

func main() {
	var bmi, tinggi float64

	fmt.Print("Masukkan BMI: ")
	fmt.Scan(&bmi)

	fmt.Print("Masukkan tinggi (m): ")
	fmt.Scan(&tinggi)

	beratIdeal := bmi * tinggi * tinggi
	fmt.Printf("Berat badan ideal adalah: %d\n", int(beratIdeal+0.5))
}