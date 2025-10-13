package main
import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3, y1, y2, y3, sisiTerpanjang float64
	fmt.Print("Masukkan nilai A: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Masukkan nilai B: ")
	fmt.Scan(&x2, &y2)
	fmt.Print("Masukkan nilai C: ")
	fmt.Scan(&x3, &y3)

	sisiAB := math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(y2 - y1, 2))
	sisiAC := math.Sqrt(math.Pow(x3 - x1, 2) + math.Pow(y3 - y1, 2))
	sisiBC := math.Sqrt(math.Pow(x3 - x2, 2) + math.Pow(y3 - y2, 2))

	sisiTerpanjang = sisiAB
	
	if sisiAC > sisiTerpanjang {
		sisiTerpanjang = sisiAC
	}
	if sisiBC > sisiTerpanjang {
		sisiTerpanjang = sisiBC
	}

	fmt.Printf("sisi terpanjangnya adalah: %.2f", sisiTerpanjang)
}