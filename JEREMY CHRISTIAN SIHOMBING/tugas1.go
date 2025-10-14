package main
import "fmt"

func main() {

    var totalBelanjaAwal, diskon int
    fmt.Print("Masukkan total belanja: ")
    fmt.Scan(&totalBelanjaAwal)
    fmt.Print("Masukkan besar diskon (%): ")
    fmt.Scan(&diskon)
    totalAkhirBelanja := totalBelanjaAwal - (totalBelanjaAwal * diskon / 100)
    fmt.Printf("Total belanja setelah diskon: %d\n", totalAkhirBelanja)
}
