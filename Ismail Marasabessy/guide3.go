package main

import "fmt"

func main() {
	var bb float64
	var tb float64
	fmt.Print("masukan bb (kg) dan masukan tb (m): ")
	fmt.Scan(&bb, &tb)
	bmi := bb / (tb * tb)
	fmt.Printf("bmi : %.2F", bmi)

}
