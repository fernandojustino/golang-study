package main

import "fmt"

func soma(nums ...int) {
	fmt.Println("nums", nums)
	total := 1
	for _, n := range nums {
		total *= n
	}
	fmt.Println(total)
	return
}

func main() {
	soma(1, 2, 3, 4, 5, 6)
}
