package main

import "fmt"

func main(){
	var bmi, tinggiBadan float64
	fmt.Print("Masukan bmi: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan Tinggi Badan: ")
	fmt.Scan(&tinggiBadan)

	berat := bmi * (tinggiBadan * tinggiBadan)
	fmt.Printf("Berat Badan: %.0f", berat)
}