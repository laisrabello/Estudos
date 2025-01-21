package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var comment string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite seu texto: ")
	comment, _ = reader.ReadString('\n')

	fmt.Print(Disemvowel(comment))
}

func Disemvowel(comment string) string {

	text := []string{}

	for i := 0; i < len(comment); i++ {
		if comment[i] != 'a' && comment[i] != 'e' && comment[i] != 'i' &&
			comment[i] != 'o' && comment[i] != 'u' && comment[i] != 'A' &&
			comment[i] != 'E' && comment[i] != 'I' && comment[i] != 'O' &&
			comment[i] != 'U' {

			text = append(text, string(comment[i]))
		}
	}
	return strings.Join(text, "")
}
