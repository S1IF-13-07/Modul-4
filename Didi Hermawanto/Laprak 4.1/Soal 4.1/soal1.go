package main

import "fmt"

func main() {
    var totalBelanja int
    var diskon int
  
    fmt.Print("Masukkan total belanja: ")
    fmt.Scan(&totalBelanja)

    fmt.Print("Masukkan diskon (%): ")
    fmt.Scan(&diskon)

    totalAkhir := float64(totalBelanja) * (1 - float64(diskon)/100)

    fmt.Println("Total belanja setelah diskon:", int(totalAkhir))
}
