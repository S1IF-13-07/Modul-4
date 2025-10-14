package main

import "fmt"

func main(){
	var totalBelanja, diskon int

	fmt.Print("Masukan total belanja: ")
	fmt.Scan(&totalBelanja)

	fmt.Print("Diskon: ")
	fmt.Scan(&diskon)

	potongan := totalBelanja * diskon / 100
	totalAkhir := totalBelanja - potongan
	fmt.Println("Total setelah diskon: ", totalAkhir)
	
}