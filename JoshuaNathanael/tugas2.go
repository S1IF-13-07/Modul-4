package main

import "fmt"

func main() {
	var BMI, TINGGIBADAN, BERATBADAN float64
	fmt.Scan(&BMI, &TINGGIBADAN)
	BERATBADAN = BMI * (TINGGIBADAN * TINGGIBADAN)
	fmt.Printf("%.0f\n", BERATBADAN)
}