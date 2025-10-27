package main

import "fmt"

func main() {
	r := 5
	const pi = 3.14
	area := pi * float64(r) * float64(r)
	fmt.Println(area)
}
