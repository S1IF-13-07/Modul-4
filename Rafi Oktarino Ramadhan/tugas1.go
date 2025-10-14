package main
import "fmt"
func main(){
    var harga, diskon, hasilharga, potongan int
    fmt.Scanln(&harga)
    fmt.Scanln(&diskon)
  
    potongan = harga * diskon/100
    hasilharga = harga - potongan
  
    fmt.Printf("%d", hasilharga)
}
