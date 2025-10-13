package main
import "fmt"
func main() {
    var hargabarang, diskon float64

    fmt.Print("Masukkan jumlah harga barang: ")
    fmt.Scan(&hargabarang)
	fmt.Print("Masukan Harga diskon: ")
	fmt.Scan(&diskon)
	potongan := hargabarang * diskon/100
	hargafiks := hargabarang - potongan
	fmt.Printf("Diskon Menjadi: %.0f", hargafiks)
}