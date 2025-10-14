package main
import "fmt"
func main(){
	var totalBelanja, diskonBelanja, totalAkhir, diskon int

	fmt.Print("Total belanja : ")
	fmt.Scanln(&totalBelanja)
	fmt.Print("Diskon yang didapat : ")
	fmt.Scanln(&diskonBelanja)

	diskon = totalBelanja * diskonBelanja/100
	totalAkhir = totalBelanja - diskon
	fmt.Println("Total belanja setelah diskon : ", totalAkhir)
}