package main

import "fmt"

func main() {
	var bmi, tinggi float64
	fmt.Print("Masukan bmi: ")
	fmt.Scanln(&bmi)
	fmt.Print("Masukan tinggi: ")
	fmt.Scanln(&tinggi)
	berat := bmi * (tinggi * tinggi)
	hasil := int(berat + 0.5)
	fmt.Print("Berat badan ideal adalah: ", hasil)
}
