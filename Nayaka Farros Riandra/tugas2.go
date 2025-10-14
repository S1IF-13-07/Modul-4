package main

import "fmt"

func main() {
	var beratbadan, tinggibadan, bmi float64

	fmt.Print("masukkan nilai Bmi = ")
	fmt.Scan(&bmi)
	fmt.Print("masukkan tinggi badan (m) = ")
	fmt.Scan(&tinggibadan)

	beratbadan = bmi * (tinggibadan * tinggibadan)
	fmt.Printf("berat badan = %.f\n", beratbadan)
}