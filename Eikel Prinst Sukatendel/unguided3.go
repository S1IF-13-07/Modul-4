package main

import "fmt"

func main() {
	var berat, tinggi float64
	fmt.Scan(&berat, &tinggi)

	bmi := berat / (tinggi * tinggi)
	fmt.Printf("%.2f\n", bmi)
}