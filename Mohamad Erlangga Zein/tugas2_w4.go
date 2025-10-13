package main
import "fmt"

func main() {
	var bmi, tinggiBadan float64
	fmt.Print("Masukkan bmi: ")
	fmt.Scan(&bmi)
	fmt.Print("masukkan tinggi badan: ")
	fmt.Scan(&tinggiBadan)

	beratBadan := bmi * (tinggiBadan * tinggiBadan)

	fmt.Printf("berat badannya adalah: %.0f", beratBadan)
}