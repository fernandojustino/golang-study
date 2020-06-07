package main

import "fmt"

func vals() (int, int) {
	return 1, 2
}

func main() {
	a, b := vals()

	_, c := vals()
	fmt.Println(a, b, c)
}
