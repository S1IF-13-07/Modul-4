package main

import "fmt"

func main() {
	var totalAwal int
	var diskon int

	fmt.Print("Masukkan total belanja: ")
	fmt.Scanln(&totalAwal)
	fmt.Print("Masukkan diskon (%): ")
	fmt.Scanln(&diskon)

	totalAkhir := totalAwal - (totalAwal * diskon / 100)

	fmt.Println("Total setelah diskon:", totalAkhir)
}