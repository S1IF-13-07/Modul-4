package main

import "fmt"

func main() {
    var bmi, tinggiBadan, beratBadan float64

    // Input dua bilangan: BMI dan tinggi badan (dalam meter)
    fmt.Printf("Masukkan BMI dan tinggi badan: ")
    fmt.Scan(&bmi, &tinggiBadan)

    // Rumus mencari berat badan
    beratBadan = bmi * (tinggiBadan * tinggiBadan)

    fmt.Printf("%.0f\n", beratBadan)
}
