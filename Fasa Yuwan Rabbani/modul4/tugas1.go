package main

import (
	"fmt"
)

func main() {
	var totalBelanja int
	var diskon int

	fmt.Print("Masukkan total belanja: ")
	fmt.Scanln(&totalBelanja)
	fmt.Print("Masukkan besar diskon (%): ")
	fmt.Scanln(&diskon)
	totalAkhir := totalBelanja - (totalBelanja * diskon / 100)
	fmt.Println("Total belanja setelah diskon:", totalAkhir)
}
