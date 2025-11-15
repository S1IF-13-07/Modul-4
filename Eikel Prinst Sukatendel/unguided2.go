package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	d1 := bilangan / 100
	d2 := (bilangan / 10) % 10
	d3 := bilangan % 10

	fmt.Println(d1 <= d2 && d2 <= d3)
}