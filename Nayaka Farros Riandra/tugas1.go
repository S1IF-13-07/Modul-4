package main

import "fmt"

func main() {
	var diskon, hargaAwal, hargaAkhir float64

	fmt.Print("Masukkan harga awal = ")
	fmt.Scan(&hargaAwal)
	fmt.Print("Masukkan diskon (%) = ")
	fmt.Scan(&diskon)

	hargaAkhir = hargaAwal - (diskon/100)*hargaAwal
	fmt.Printf("Harga setelah diskon = %.2f", hargaAkhir)
}