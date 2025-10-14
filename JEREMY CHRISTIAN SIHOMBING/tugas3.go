package main
import (
    "fmt"
    "math"
)
func main() {
    var x1, y1, x2, y2, x3, y3 float64
    fmt.Print("Masukkan koordinat titik A (x y): ")
    fmt.Scan(&x1, &y1)
    fmt.Print("Masukkan koordinat titik B (x y): ")
    fmt.Scan(&x2, &y2)
    fmt.Print("Masukkan koordinat titik C (x y): ")
    fmt.Scan(&x3, &y3)
    sisiab := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
    sisibc := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
    sisica := math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2))
    Terpanjang := sisiab
    if sisibc > Terpanjang {
        Terpanjang = sisibc
    }
    if sisica > Terpanjang {
        Terpanjang = sisica
    }
    fmt.Printf("Sisi terpanjang: %.2f\n", Terpanjang)
}
