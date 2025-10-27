package main

import "fmt"

func factorial(n int) int {
	k := 1
	for i := 1; i <= n; i++ {
		k *= i
	}
	return k
}

func main() {
	n := 10
	fmt.Println(factorial(n))
}
