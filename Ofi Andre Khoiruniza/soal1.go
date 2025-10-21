
package main

import "fmt"

func main () {
	var totalBelanja, diskon float64
	fmt.Print("Input, Total Belanja : ")
	fmt.Scanln(&totalBelanja)
	fmt.Print("Input, persen diskon : ")
	fmt.Scanln(&diskon)
	fmt.Printf("Output, Total Belanja : %.0f", totalBelanja-(totalBelanja*diskon/100))
}
