package main
import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Print("Masukkan berat badan: ")
	fmt.Scan(&beratBadan)
	fmt.Print("Masukkan tinggi badan: ")
	fmt.Scan(&tinggiBadan)

	bmi = beratBadan / (tinggiBadan * tinggiBadan)
	fmt.Printf("bmi nya adalah: %.2f", bmi)
}