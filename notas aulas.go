package main //marca o inicio de um programa

import (
	"fmt"
	//"time"
	// mostra como renomear um pacote
	//Lais "strings"
)

func main() {
	/*fmt.Println("Hello, world!")

	//chama o pacote strings renomeado para Lais
	fmt.Println(Lais.Split("Laís Rabello", ""))

	//Como declarar variáveis e constantes
	var nome string
	nome = "Laís"

	fmt.Println(nome)

	//declara a variável sem precisar do tipo
	var nome1 = "Lais"
	fmt.Println(nome1)

	//declara a variável sem precisar do var e do tipo
	nome2 := "Laís Rabello"

	fmt.Println(nome2)

	//declarar constantes
	const idade_lais = 27

	fmt.Println(idade_lais)

	//Zero Values: quando declaramos uma variável e não
	//atribuimos valor a ela, o Go atribui um valor defaul padrão da língua
	//Inteiro: 0
	//Float: 0
	//Bool: False
	//String: ""
	//----------------------------------------------------

	//F U N Ç Õ E S

	//chamando a função soma fora da função main
	fmt.Println("A soma é:", soma(10, 17))
	fmt.Println("O Login é:", Login("Lasa_Wally"))*/

	//-----------------------------------------------------

	//LISTAS

	/*1. Arrays e Slices: listas Homogêneas
		-> todos os elementos têm o mesmo tipo
		exemplo: [1, 2, 3, 4, 5] - int []

		Array: Tamanho fixo
		Slice: Sem tamanho fixo


	  2. Maps: Heterogêneos
	    -> pode misturar tipos
		-> modelo [chave:valor] */

	//declara uma variável e guarda um array
	/*var array [2]string
	array[0] = "Hello"
	array[1] = "Bitch"

	fmt.Println(array)
	fmt.Println(array[1])

	arrayNum := [5]int{7, 1997, 11, 2, 4}
	fmt.Println(arrayNum)
	fmt.Println(arrayNum[0:3])

	//declarar uma variável e guarda um slice
	slice := make([]string, 5)
	slice[0] = "Hello"
	slice[1] = "Fetuccini"
	fmt.Println(slice)
	fmt.Println(slice[1])

	slice[2] = "Laís"
	fmt.Println(slice)

	fmt.Println(len(slice))

	slice = append(slice, "Legal", "Que", "Dia", "Legal")
	fmt.Println(slice)
	fmt.Println("-------------------------------------------")

	//declarar um map e guardar em uma variável
	idade := map[string]int{}
	idade["Laís"] = 27
	idade["Vinicius"] = 27

	fmt.Println(idade)
	fmt.Println(idade["Laís"])
	fmt.Println(idade["Vinicius"])

	//outra forma de criar um map
	anoNasc := map[string]int{
		"Laís":     1997,
		"Vinicius": 1997,
	}

	fmt.Println(anoNasc)

	//adicionando outro elemento ao map
	anoNasc["Laís Code"] = 2024
	fmt.Println(anoNasc)

	//---------------------------------------------------------

	//STRUCTS

	//criando a estrutura do tipo Pessoa
	type Pessoa struct {
		Nome      string
		Idade     int
		Profissao string
		Salario   float64
		Casada    bool
	}

	//utilizando a estrutura do tipo Pessoa
	fmt.Println("--------------------------------------------------------")
	fmt.Println(Pessoa{"Laís", 27, "Física", 4500.80, false})
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	//----------------------------------------------------------------------
	//C O N D I C I O N A I S IF/ELSE

	var age int
	fmt.Println("------------------------------------------------------")
	fmt.Println()

	fmt.Println("Descubra se pode ingerir bebida alcóolica -->")
	fmt.Println("Digite sua idade:")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("Proíbido bebida alcóolica para você!")
	} else {
		fmt.Println("TA LIBERADOOOO lol")
	}

	fmt.Println()
	fmt.Println("------------------------------------------------------")

	//----------------------------------------------------------
	//C O N D I C I O N A I S SWITCH CASE

	name := "Spike"

	switch name {
	case "Laís":
		fmt.Println("É uma pessoa")
	case "Bob":
		fmt.Println("É um personagem")
	case "Spike":
		fmt.Println("É um cachorro")
	}*/

	//------------------------------------------------------------------
	// L O O P S

	//Repetir tarefas

	/*sum := 0

	for i := 0; i < 10; i++ {
		fmt.Println("Sum: ", sum)
		fmt.Println("Indice: ", i)
		sum += i
	}
	fmt.Println(sum)*/

	// loop infinito
	/*
		for {
			fmt.Println("Golang do zero !!!")
			time.Sleep(2 * time.Second)
		}*/

	// for range

	frutas := []string{"laranja", "maçã", "banana", "uva", "kiwi"}
	for i, fruta := range frutas {
		fmt.Println("Fruta", fruta, "Indice", i)
	}

}

// Declarar F U N Ç Õ E S

// função começando com letra minúscula > PRIVADA
// quer dizer que só pode ser chamada em um pacote (nesse caso, o main)
func soma(x int, y int) int {
	return x + y
}

// função começando com letra maiúscula > PÚBLICA
// quer dizer que pode ser chamada em mais de um pacote, ou seja, dentro e fora do pacote main
func Login(nome string) string {
	return nome

}
