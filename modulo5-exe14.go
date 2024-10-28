package main

import (
	"fmt"
	"strings"
)

func main() {

	a := []string{"az", "toto", "picaro", "zone", "kiwi"}

	fmt.Println(PartList(a))
}

func PartList(arr []string) string {
	var result string

	for i := 1; i < len(arr); i++ {

		word := fmt.Sprintf("(%s, %s)", strings.Join(arr[:i], " "), strings.Join(arr[i:], " "))

		result += word
	}
	return result
}
