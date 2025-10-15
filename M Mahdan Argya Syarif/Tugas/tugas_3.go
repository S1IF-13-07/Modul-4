package main

import (
    "fmt"
    "math"
)

func Hasil(ax float64, ay float64, bx float64, by float64, cx float64, cy float64) float64 {
    AB := math.Sqrt(math.Pow((bx-ax), 2) + math.Pow((by-ay), 2))
    BC := math.Sqrt(math.Pow((cx-bx), 2) + math.Pow((cy-by), 2))
    CA := math.Sqrt(math.Pow((ax-cx), 2) + math.Pow((ay-cy), 2))
    longest := AB
    if BC > longest {
        longest = BC
    }
    if CA > longest {
        longest = CA
    }
    return longest
}

func main() {
    var ax, ay, bx, by, cx, cy float64
    fmt.Print("Masukkan koordinat titik A (x y): ")
    fmt.Scanln(&ax, &ay)
    fmt.Print("Masukkan koordinat titik B (x y): ")
    fmt.Scanln(&bx, &by)
    fmt.Print("Masukkan koordinat titik C (x y): ")
    fmt.Scanln(&cx, &cy)
    hasil := Hasil(ax, ay, bx, by, cx, cy)
    fmt.Printf("Panjang sisi terpanjang adalah: %.2f\n", hasil)
}
