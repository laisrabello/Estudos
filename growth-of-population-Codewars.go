package main

import "fmt"

func main() {
	p0 := 1500000
	percent := 0.02
	aug := -1000
	p := 2000000

	fmt.Print(NbYear(p0, percent, aug, p))
}

func NbYear(p0 int, percent float64, aug int, p int) int {

	cont := 0
	p0Float := float64(p0)
	autFloat := float64(aug)
	pFloat := float64(p)

	for i := 0; p0Float < pFloat; i++ {

		p0Float += (p0Float * percent) + autFloat

		cont++

	}
	return cont
}
