package main

import "fmt"

func main() {
	var bmi, tinggi float64
	
	fmt.Scan (&bmi, &tinggi)
	berat := bmi * (tinggi * tinggi)

	fmt.Printf("%.f\n", berat)
}