package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var caracter string

	fmt.Scan(&caracter)

	inteiro, _ := strconv.Atoi(caracter)

	fmt.Printf("O caracter '%s' do tipo %s foi convertido para o inteiro '%d' do tipo %s\n",
		caracter, reflect.TypeOf(caracter), inteiro, reflect.TypeOf(inteiro))

}
