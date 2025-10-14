package main

import "fmt"

func main() {
    var totalBelanja, diskonPersen float64
    var totalAkhir float64

    fmt.Print("Masukkan total belanja: ")
    fmt.Scan(&totalBelanja)
    fmt.Print("Masukkan besarnya diskon (%): ")
    fmt.Scan(&diskonPersen)

    totalAkhir = totalBelanja - (totalBelanja * diskonPersen / 100)
    fmt.Printf("Total belanja setelah diskon: %.0f\n", totalAkhir)
}
