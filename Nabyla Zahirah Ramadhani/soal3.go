package main

import (
	"fmt"
	"math"
)

func main (){
	var ax, ay, bx, by, cx, cy float64

	fmt.Print("Masukkan koordinat titik ax: ")
	fmt.Scan(&ax)
	fmt.Print("Masukkan koordinat titik ay: ")
	fmt.Scan(&ay)
	fmt.Print("Masukkan koordinat titik bx: ")
	fmt.Scan(&bx)
	fmt.Print("Masukkan koordinat titik by: ")
	fmt.Scan(&by)
	fmt.Print("Masukkan koordinat titik cx: ")
	fmt.Scan(&cx)
	fmt.Print("Masukkan koordinat titik cy: ")
	fmt.Scan(&cy)

	ab := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	maxSide := ab
	if bc > maxSide {
		maxSide = bc
	}
	if ca > maxSide {
		maxSide = ca
	}
	fmt.Printf("Sisi terpanjang: %.2f\n", maxSide)


}