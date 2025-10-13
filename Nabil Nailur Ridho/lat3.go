package main

import (
	"fmt"
	"math"
)

func main(){
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Print("Masukan nilai A: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Masukan nilai B: ")
	fmt.Scan(&x2, &y2)
	fmt.Print("Masukan nilai C: ")
	fmt.Scan(&x3, &y3)
	sisiAB := math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(y2 - y1, 2))
	sisiAC := math.Sqrt(math.Pow(x3 - x1, 2) + math.Pow(y3 - y1, 2))
	sisiBC := math.Sqrt(math.Pow(x3 - x2, 2) + math.Pow(y3 - y2, 2))

	sisiTerpanjang := sisiAB
	if sisiAC > sisiTerpanjang {
		sisiTerpanjang = sisiAC
	} else if sisiBC > sisiTerpanjang {
		sisiTerpanjang = sisiBC
	}
	fmt.Printf("Sisi Terpanjang Adalah: %.0f", sisiTerpanjang)
}