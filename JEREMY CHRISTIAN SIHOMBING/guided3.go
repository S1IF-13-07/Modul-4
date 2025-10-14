package main
import "fmt"

func main() {

 var beratBadan, tinggiBadan, bmi float64
 fmt.Print("Masukkan berat badan (kg) dan tinggi badan (m): ")
 fmt.Scan(&beratBadan, &tinggiBadan)
 bmi = beratBadan / (tinggiBadan * tinggiBadan)
 fmt.Printf("Nilai BMI: %.2f\n", bmi)
}
