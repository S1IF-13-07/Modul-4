package main

import (
    "fmt"
)

func Bmi(berat float64, tinggi float64) float64 {
    if tinggi <= 0 {
        fmt.Println("Tinggi harus lebih besar dari 0")
        return 0
    }
    hasil := berat / (tinggi * tinggi)
    return hasil
}

func main() {
    var height, weight float64
    fmt.Print("Masukkan berat (kg): ")
    fmt.Scanln(&weight)
    fmt.Print("Masukkan tinggi (m): ")
    fmt.Scanln(&height)
    bmi := Bmi(weight, height)
    fmt.Printf("BMI Anda adalah: %.0f\n", bmi)
}
