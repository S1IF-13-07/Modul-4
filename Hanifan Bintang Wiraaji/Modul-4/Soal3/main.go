package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)
	var sab = math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(y2 - y1, 2))
	var sbc = math.Sqrt(math.Pow(x3 - x2, 2) + math.Pow(y3 - y2, 2))
	var sca = math.Sqrt(math.Pow(x1 - x3, 2) + math.Pow(y1 - y3, 2))
	var terpanjang float64 = sab
	if sbc > terpanjang {
		terpanjang = sbc
	}
	if sca > terpanjang {
		terpanjang = sca
	}
	fmt.Printf("Sisi terpanjang: %.2f", terpanjang)
}