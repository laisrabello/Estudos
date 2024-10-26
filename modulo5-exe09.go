package main

//Função que diz se o número inteiro dado é Par ou Ímpar
//Function that says if an integer is Even or Odd

func EvenOrOdd(n int) string {
	if n%2 == 0 {
		return "Even"
	}

	return "Odd"

}

//Solução diferente utilizando operação bit a bit

/*func EvenOrOdd(number int) string {
	return []string{"Even", "Odd"}[number & 1]
}*/
