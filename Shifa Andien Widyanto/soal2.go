package main
import "fmt"
func main(){
	var bmi, tb, bb float64
	fmt.Print("Masukan nilai BMI : ")
	fmt.Scanln(&bmi)
	fmt.Print("Tinggi badan : ")
	fmt.Scanln(&tb)

	bb = bmi * (tb * tb)
	fmt.Printf("Berat badannya : %0.f\n", bb)
}