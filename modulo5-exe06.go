package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var inteiro int

	fmt.Scan(&inteiro)

	caracter := strconv.Itoa(inteiro)

	fmt.Printf("O inteiro '%d' do tipo %s foi convertido para o caracter '%s' do tipo %s\n",
		inteiro, reflect.TypeOf(inteiro), caracter, reflect.TypeOf(caracter))

}
