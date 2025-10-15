package main

import (
	"fmt"
	"math"
)

func main() {
	var bmi, tinggi float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan (meter): ")
	fmt.Scan(&tinggi)

	// Rumus berat = BMI * (tinggi^2)
	berat := bmi * math.Pow(tinggi, 2)

	// Dibulatkan ke bilangan bulat terdekat
	beratBulat := int(math.Round(berat))

	fmt.Printf("Berat badan: %d kg\n", beratBulat)
}
