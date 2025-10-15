package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64


	fmt.Scan(&ax, &ay)
	fmt.Scan(&bx, &by)
	fmt.Scan(&cx, &cy)


	AB := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	BC := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	CA := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	
	terpanjang := AB
	if BC > terpanjang {
		terpanjang = BC
	}
	if CA > terpanjang {
		terpanjang = CA
	}


	fmt.Printf("%.2f\n", terpanjang)
}
