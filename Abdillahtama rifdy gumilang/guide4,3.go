package main

import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Print("Masukkan berat badanmu kawan: ")
	fmt.Scan(&beratBadan)
	fmt.Print("Masukkan tinggi badan mu kink: ")
	fmt.Scan(&tinggiBadan)
	bmi = beratBadan / (tinggiBadan * tinggiBadan)
	fmt.Printf("hasil bmi nya adalah: %.2f", bmi)
}
