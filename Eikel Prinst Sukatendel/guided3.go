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

	sisiAB := math.Hypot(x2-x1, y2-y1)
	sisiBC := math.Hypot(x3-x2, y3-y2)
	sisiCA := math.Hypot(x1-x3, y1-y3)

	terpanjang := sisiAB
	if sisiBC > terpanjang {
		terpanjang = sisiBC
	}
	if sisiCA > terpanjang {
		terpanjang = sisiCA
	}

	fmt.Printf("%.2f\n", terpanjang)
}