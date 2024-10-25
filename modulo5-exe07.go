package main

import "fmt"

func main() {

	var caracter string
	var number int

	fmt.Println("Digite a palavra para repetir:")
	fmt.Scan(&caracter)
	fmt.Println("Digite o número de repetições: ")
	fmt.Scan(&number)

	fmt.Printf(repeat(number, caracter))

}

func repeat(n int, s string) string {
	var resultado string

	for i := 0; i < n; i++ {
		resultado += s
	}
	return resultado
}
