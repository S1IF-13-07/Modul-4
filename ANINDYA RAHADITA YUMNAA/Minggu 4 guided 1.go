package main

import "fmt"

func main () {
	var totalbelanja, diskon int
	fmt.Scan(&totalbelanja)
	fmt.Scan(&diskon)

	totalakhir := totalbelanja - (totalbelanja * diskon / 100)
	fmt.Println(totalakhir)
}