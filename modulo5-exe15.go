package main

import (
	"fmt"
)

func main() {
	s := "Codewars"

	fmt.Println(SortMyString(s))
}

func SortMyString(s string) string {
	var even string
	var odd string

	for i := 0; i < len(s); i++ {

		if i%2 == 0 {
			even += string(s[i])
		} else {
			odd += string(s[i])
		}
	}

	return fmt.Sprintf("%s %s", even, odd)
}
