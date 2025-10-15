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
	AB := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	BC := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	CA := math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2))
	terpanjang := AB
	if BC > terpanjang {
		terpanjang = BC
	}
	if CA > terpanjang {
		terpanjang = CA
	}
	fmt.Print(terpanjang)
}
