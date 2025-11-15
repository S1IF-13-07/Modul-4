package main

import "fmt"

func main() {
	var detik int
	fmt.Scan(&detik)

	jam := detik / 3600
	menit := (detik % 3600) / 60
	sisaDetik := detik % 60

	fmt.Printf("%d jam %d menit dan %d detik\n", jam, menit, sisaDetik)
}