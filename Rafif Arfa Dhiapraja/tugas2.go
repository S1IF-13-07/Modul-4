package main

import "fmt"

func main() {
	var beratbadan, tinggibadan, bmi float64

	fmt.Print("Masukkan BMI dan Tinggi Badan (m) : ")
	fmt.Scan(&bmi, &tinggibadan)

	beratbadan = bmi * tinggibadan * tinggibadan
	fmt.Printf("%.0f kg", beratbadan)
}