package main

import "fmt"

func main () {
	var nilaiBMI, tinggi float64
	fmt.Scan(&nilaiBMI, &tinggi)
	berat := nilaiBMI * (tinggi * tinggi)
	fmt.Printf("%.f\n", berat)

}