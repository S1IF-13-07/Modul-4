package main

import "fmt"

func main() {
	var totalBelanja, diskon int
	var totalAkhir float64

	fmt.Scan(&totalBelanja)
	fmt.Scan(&diskon)
	 
	totalAkhir = float64(totalBelanja) * (1 - float64(diskon)/100)

	fmt.Printf("Total belanja akhir: %.2f\n", totalAkhir)
}