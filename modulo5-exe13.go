package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	games := []string{"1:0", "2:0", "3:3", "4:0", "2:1",
		"1:3", "1:4", "2:3", "2:4", "3:3"}

	fmt.Println(Points(games))
}

func Points(games []string) int {

	totalPoints := 0

	for _, game := range games {
		scores := strings.Split(game, ":")

		homeScore, _ := strconv.Atoi(scores[0])
		awayScore, _ := strconv.Atoi(scores[1])

		if homeScore > awayScore {
			totalPoints += 3
		} else if homeScore == awayScore {
			totalPoints += 1
		}
	}
	return totalPoints

}

// OUTRA SOLUÇÃO

/*
package kata

import ( "strings" )

func Points(games []string) int {
  result := 0
  for _, game := range games {
    str := strings.Split(game, ":")
    x, y := str[0], str[1]
    switch {
      case x > y:
        result += 3
      case x == y:
        result += 1
    }
  }
  return result
}
*/

/*package kata

func Points(games []string) (res int) {
  for _,v:=range games {
  if v[0]>v[2] {res +=3}
  if v[0]==v[2] {res +=1}
}
  return
}
*/

/*
package kata

const WIN_POINTS = 3
const DRAW_POINTS = 1

func Points(games []string) int {
  var totalPoints int
  for _, val := range games {
    x := string(val[:1])
    y := string(val[2:])

    if (x > y) {
      totalPoints += WIN_POINTS
    } else if (x == y) {
      totalPoints += DRAW_POINTS
    }
  }
  return totalPoints
}
*/
