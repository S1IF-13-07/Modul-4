package main
import "fmt"

func main() {

    var bmi, tinggi, berat float64
    fmt.Print("Masukkan nilai BMI dan tinggi badan (meter): ")
    fmt.Scan(&bmi, &tinggi)
    berat = bmi * (tinggi * tinggi)
    fmt.Printf("Berat badan: %.0f kg\n", berat)
}
