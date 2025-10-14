package main

import (
	"fmt"
	"math"

)

func main() {
	var xA, yA, xB, yB, xC, yC float64
	var AB, BC, CA, maxSide float64

	// Input koordinat titik A
	fmt.Print("Masukkan koordinat titik A (x y): ")
	fmt.Scan(&xA, &yA)

	// Input koordinat titik B
	fmt.Print("Masukkan koordinat titik B (x y): ")
	fmt.Scan(&xB, &yB)

	// Input koordinat titik C
	fmt.Print("Masukkan koordinat titik C (x y): ")
	fmt.Scan(&xC, &yC)

	// Hitung panjang sisi-sisi segitiga
	AB = math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
	BC = math.Sqrt(math.Pow(xC-xB, 2) + math.Pow(yC-yB, 2))
	CA = math.Sqrt(math.Pow(xA-xC, 2) + math.Pow(yA-yC, 2))

	// Tentukan sisi terpanjang
	maxSide = AB
	if BC > maxSide {
		maxSide = BC
	}
	if CA > maxSide {
		maxSide = CA
	}

	// Tampilkan hasil
	fmt.Printf("Panjang sisi terpanjang: %.2f\n", maxSide)
}
