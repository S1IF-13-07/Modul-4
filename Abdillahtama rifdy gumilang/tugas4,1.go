package main

import "fmt"

func main() {
	var hargaawal, diskon int
	fmt.Print("Masukan harga awal: ")
	fmt.Scanln(&hargaawal)
	fmt.Print("Masukan diskon: ")
	fmt.Scanln(&diskon)
	hasil := hargaawal - (hargaawal * diskon / 100)
	fmt.Println("Harga setelah diskon:", hasil)
}
