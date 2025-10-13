package main
import "fmt"

func main(){
	var hargaAwal, persentaseDiskon int
	fmt.Print("masukkan harga awal: ")
	fmt.Scan(&hargaAwal)
	fmt.Print("masukkan persentase diskon: ")
	fmt.Scan(&persentaseDiskon)

	potongan := hargaAwal * persentaseDiskon / 100
	hargaFix := hargaAwal - potongan

	fmt.Print("Harga setelah diskon: ", hargaFix)
}