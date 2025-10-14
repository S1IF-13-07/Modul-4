package main
import (
 "fmt"
 "math"
)
func main() {

	var bmi, tinggibadan float64

	fmt.Scanln(&bmi, &tinggibadan)

	berat := bmi * (tinggibadan * tinggibadan)

	fmt.Println(int(math.Round(berat)))
}
