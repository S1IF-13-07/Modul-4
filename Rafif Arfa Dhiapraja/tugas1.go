package main

import "fmt"

func main() {
	var harga int
	fmt.Scan(&harga)

	var diskon int
	fmt.Scan(&diskon)

	var total int
	total = harga - (harga * diskon / 100)
	fmt.Print(total)
}