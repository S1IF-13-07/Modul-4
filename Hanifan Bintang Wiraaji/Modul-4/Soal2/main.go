package main

import "fmt"

func main() {
	var bmi, tb float64
	fmt.Println("Masukan BMI dan Tinggi Badan(Dalam Meter):")
	fmt.Scanln(&bmi, &tb)
	var bb float64 = ( ( tb * tb ) * bmi ) + 0.1
	var bbbulat int = int(bb) 
	fmt.Printf("Berat badan(Dalam KG): %v", bbbulat)
}