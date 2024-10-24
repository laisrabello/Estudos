package main

import "fmt"

func main() {

	inteiros := []int{}

	var sum int

	for _, inteiro := range inteiros {
		if inteiro > 0 {
			sum += inteiro
		} else {
			sum = sum
		}
	}
	fmt.Println(sum)
}
