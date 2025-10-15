package main
import "fmt"
func main () {
	var totalBelanja, diskon int 
	fmt.Print("Masukan total belanja: ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Masukan diskon (%): ")
	fmt.Scan(&diskon)
	totalAkhir:= totalBelanja - (totalBelanja * diskon/100)
	fmt.Println("TOTAL SEMUA BELANJAAN:",totalAkhir)

}