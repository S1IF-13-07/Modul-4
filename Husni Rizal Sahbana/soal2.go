package main

import "fmt"

func main() {
	var bmi, tinggiBadan float64

	fmt.Print("Masukan BMI kamu : ")
	fmt.Scan(&bmi)

	fmt.Print("Masukan tinggi badan : ")
	fmt.Scan(&tinggiBadan)

	berat := bmi * (tinggiBadan * tinggiBadan)
	fmt.Printf("%.0f", berat)

}
