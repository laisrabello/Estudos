//E N U N C I A D O
/*Simple, given a string of words, return the length of the shortest
word(s). String will never be empty and you do not need to account for
different data types.*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "bitcoin take over we the world maybe who knows perhaps"

	words := strings.Split(s, " ")
	length := words[0]

	for i := 1; i < len(words); i++ {

		if len(words[i]) < len(length) {
			length = words[i]

		}

	}

	fmt.Println(len(length))

}

/*func FindShort(s string) int {
	words := strings.Split(s, " ")
	var length int

	for i, word := range words {

		length = len(words[0])

		if len(words[i]) < length {
			length = len(word)
		}

	}
	return length
}*/
