package main

import "fmt"

func main() {
	var total_belanja, diskon int

	fmt.Print("Masukan total belanja mu : ")
	fmt.Scan(&total_belanja)

	fmt.Print("Diskon yang anda peroleh : ")
	fmt.Scan(&diskon)

	potongan := total_belanja * diskon / 100
	total_akhir := total_belanja - potongan

	fmt.Println("Jadi total akhir :  ", total_akhir)
}
