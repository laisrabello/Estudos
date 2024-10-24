package main

import "fmt"

func main() {

	var num int
	var sum int

	fmt.Scan(&num)

	for i := 1; i < num+1; i++ {

		sum += i
	}

	fmt.Println(sum)
}
