package main

import (
	"fmt"
	"math"
)

func main() {
	var xA, yA, xB, yB, xC, yC float64

	fmt.Print("Masukkan koordinat titik A (x) : ")
	fmt.Scan(&xA)

	fmt.Print("Masukkan koordinat titik A (y) : ")
	fmt.Scan(&yA)

	fmt.Print("Masukkan koordinat titik B (x) : ")
	fmt.Scan(&xB)

	fmt.Print("Masukkan koordinat titik B (y) : ")
	fmt.Scan(&yB)

	fmt.Print("Masukkan koordinat titik C (x): ")
	fmt.Scan(&xC)

	fmt.Print("Masukkan koordinat titik C (y): ")
	fmt.Scan(&yC)

	AB := math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
	BC := math.Sqrt(math.Pow(xC-xB, 2) + math.Pow(yC-yB, 2))
	CA := math.Sqrt(math.Pow(xA-xC, 2) + math.Pow(yA-yC, 2))

	maxSide := AB
	if BC > maxSide {
		maxSide = BC
	}
	if CA > maxSide {
		maxSide = CA
	}

	fmt.Printf("Sisi terpanjang adalah: %.2f\n", maxSide)
}
