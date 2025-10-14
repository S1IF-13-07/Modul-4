package main

import "fmt"

func main() {
	var harga, diskon int
	fmt.Print("Masukan harga: ")
	fmt.Scanln(&harga)
	fmt.Print("Masukan diskon: ")
	fmt.Scanln(&diskon)
	var hasil int = harga - ( harga * diskon / 100 )
	fmt.Printf("Harga total: %v", hasil)
}