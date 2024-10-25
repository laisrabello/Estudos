package main

import "fmt"

func main() {

	arrays := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	retorno := [2]int{}
	cont := 0
	var sum int

	for _, array := range arrays {
		if array > 0 {
			cont += 1

		} else {
			sum += array
		}
	}
	retorno[0] = cont
	retorno[1] = sum

	fmt.Println(retorno)
}
