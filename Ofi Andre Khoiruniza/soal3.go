package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3, y1, y2, y3, sisi1, sisi2, sisi3 float64
	fmt.Println("Input, 3 Titik Koordinat Segitiga (x,y) : ")
	fmt.Scan(&x1)
	fmt.Scanln(&y1)
	fmt.Scan(&x2)
	fmt.Scanln(&y2)
	fmt.Scan(&x3)
	fmt.Scanln(&y3)
	sisi1 = math.Sqrt((math.Pow(x1-x2, 2)) + (math.Pow(y1-y2, 2)))
	sisi2 = math.Sqrt((math.Pow(x2-x3, 2)) + (math.Pow(y2-y3, 2)))
	sisi3 = math.Sqrt((math.Pow(x1-x3, 2)) + (math.Pow(y1-y3, 2)))
	fmt.Printf("%.2f", math.Max(sisi1, math.Max(sisi2, sisi3)))
}