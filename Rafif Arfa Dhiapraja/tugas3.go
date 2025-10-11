package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay float64
	var bx, by float64
	var cx, cy float64

	
	fmt.Scan(&ax, &ay)
	fmt.Scan(&bx, &by)
	fmt.Scan(&cx, &cy)

	
	ab := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	
	if ab >= bc && ab >= ca {
		fmt.Printf("%.2f",ab)
		} else if bc >= ab && bc >= ca {
			fmt.Printf("%.2f", bc)
		} else {
			fmt.Printf("%.2f", ca)
		}

}
