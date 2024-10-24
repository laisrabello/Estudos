package main

import "fmt"

func main() {

	var number int

	fmt.Println("Digite um número:")
	fmt.Scanln(&number)

	if number > 0 {
		num := number * -1
		fmt.Println("O negativo deste número é: ", num)
	} else {

		fmt.Println("Nada a ser feito")

	}

}
