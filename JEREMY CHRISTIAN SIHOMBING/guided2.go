package main
import "fmt"

func main() {

 var bilangan, d1, d2, d3 int
 fmt.Print("Masukkan bilangan 3 digit: ")
 fmt.Scan(&bilangan)
 d1 = bilangan / 100
 d2 = bilangan % 100 / 10
 d3 = bilangan % 100 % 10
 fmt.Printf("Apakah bilangan terurut membesar atau tidak? %v\n", d1 <= d2 && d2 <= d3)
}
