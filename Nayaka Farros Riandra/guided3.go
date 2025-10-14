package main

import "fmt"

func main() {
	var beratbadan, tinggibadan, bmi float64

	fmt.Print("masukkan berat badan (kg) = ")
	fmt.Scan(&beratbadan)
	fmt.Print("masukkan tinggi badan (m) = ")
	fmt.Scan(&tinggibadan)

	bmi = beratbadan / (tinggibadan * tinggibadan)
	
	fmt.Printf("BMI = %.2f\n", bmi)
}