package main

import "fmt"

func soma(a int, b int) int {
	return (a * b)
}

func main() {
	for i := 1; i <= 10; i++ {
		for a := 1; a <= 10; a++ {
			fmt.Println(soma(i, a))
		}
		fmt.Println("-----------------------------")
	}
}
