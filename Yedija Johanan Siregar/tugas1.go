package main

import "fmt"

func main() {
    var totalBelanja, diskon int
    fmt.Printf("Masukkan total belanja: ")
    fmt.Scan(&totalBelanja)
    fmt.Printf("Masukkan persentase diskon: ")
    fmt.Scan(&diskon)

    totalAkhir := float64(totalBelanja) - (float64(totalBelanja) * float64(diskon) / 100)

    fmt.Printf("%.0f\n", totalAkhir)
}
