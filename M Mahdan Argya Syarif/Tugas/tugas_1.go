package main

import (
    "fmt"
)

func Discount(price int, percentage int) int {
    if percentage < 0 || percentage > 100 {
        fmt.Println("Percentage must be between 0 and 100")
        return price
    }
    discountAmount := (price * percentage) / 100
    return price - discountAmount
}

func main() {
    var harga, persen int
    fmt.Print("Masukkan harga barang: ")
    fmt.Scanln(&harga)
    fmt.Print("Masukkan persentase diskon: ")
    fmt.Scanln(&persen)
    hargaSetelahDiskon := Discount(harga, persen)
    fmt.Printf("Harga setelah diskon: %d\n", hargaSetelahDiskon)
}
