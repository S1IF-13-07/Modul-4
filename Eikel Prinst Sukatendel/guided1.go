package main

import "fmt"

func main() {
	var hargaAwal, diskon int

	fmt.Print("Masukkan harga awal: ")
	fmt.Scan(&hargaAwal)

	fmt.Print("Masukkan diskon (%): ")
	fmt.Scan(&diskon)

	hargaAkhir := hargaAwal * (100 - diskon) / 100
	fmt.Printf("Harga setelah diskon: %d\n", hargaAkhir)
}