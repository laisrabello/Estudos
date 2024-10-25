package main

import "fmt"

func main() {

	exemplo := []int{1, -2, -10, -4, 5, 8}

	fmt.Println(inverse(exemplo))
}

func inverse(arrays []int) []int {
	resultado := make([]int, len(arrays))

	for i, array := range arrays {
		resultado[i] = array * (-1)
	}
	return resultado
}
