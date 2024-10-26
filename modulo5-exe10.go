package main

import "fmt"

func main() {
	n := 5

	fmt.Println(CalculateYears(n))
}

func CalculateYears(humanYears int) (result [3]int) {

	catYears, dogYears := 15, 15

	switch {

	case humanYears == 1:
		catYears, dogYears = 15, 15

	case humanYears == 2:
		catYears, dogYears = 24, 24

	case humanYears > 2:
		catYears, dogYears = 24, 24

		for i := 3; i < humanYears+1; i++ {
			catYears += 4
			dogYears += 5
		}
	}

	result[0] = humanYears
	result[1] = catYears
	result[2] = dogYears

	return result
}
