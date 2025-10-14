package main

import "fmt"

func main (){
	var bmi, tinggiBadan float64

	fmt.Print("Masukkan BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan: ")
	fmt.Scan(&tinggiBadan)
	berat := bmi * tinggiBadan * tinggiBadan
	fmt.Printf("Berat badan: %.0f", berat)
}