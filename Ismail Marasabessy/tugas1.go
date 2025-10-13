package main

import "fmt"

func main() {
	var totalBelanja, diskon int
	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Masukkan diskon (%): ")
	fmt.Scan(&diskon)
	totalAkhir := totalBelanja - (totalBelanja * diskon / 100)
	fmt.Println("Total belanja akhir:", totalAkhir)
}
