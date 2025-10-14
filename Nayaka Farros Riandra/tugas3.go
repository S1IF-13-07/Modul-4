package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3, y1, y2, y3 float64
	
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	fmt.Scan(&x3)
	fmt.Scan(&y3)
	
	a := math.Sqrt(math.Pow((x2 - x1), 2) + math.Pow((y2 - y1), 2))
	b := math.Sqrt(math.Pow((x3 - x2), 2) + math.Pow((y3 - y2), 2))
	c := math.Sqrt(math.Pow((x1 - x3), 2) + math.Pow((y1 - y3), 2))

	terpanjang := a
	if b > terpanjang {
		terpanjang = b
	}
	if c > terpanjang {
		terpanjang = c
	}	
	fmt.Printf("%.2f\n", terpanjang)
}