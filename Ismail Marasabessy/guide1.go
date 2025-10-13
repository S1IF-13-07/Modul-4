package main

import "fmt"

func main() {
	var detik, menit, jam int
	fmt.Print("jumlah detik: ")
	fmt.Scanln(&detik)

	jam = detik / 3600
	menit = detik % 60
	detik = detik % 60

	fmt.Printf("%d jam, %d menit, %d detik", jam, menit, detik)
}
