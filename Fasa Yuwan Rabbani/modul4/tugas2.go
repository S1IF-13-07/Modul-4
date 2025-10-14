package main

import (
	"fmt"
)

func main() {
	var bmi, tinggi float32

	fmt.Print("Masukkan nilai BMI = ")
	fmt.Scanln(&bmi)
	fmt.Print("Masukkan tinggi badan = ")
	fmt.Scanln(&tinggi)
	berat := bmi * (tinggi * tinggi)
	fmt.Printf("Berat badan anda adalah %.2f", berat)
}
