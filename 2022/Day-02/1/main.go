package main

import (
	"fmt"
	"os"
	"strings"
)

type game [][]string

func main() {

	content, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
	games := CreatingPairs(string(content))

	fmt.Println(games)

	var sum int
	for _, g := range games {
		opponent := g[0]
		me := g[1]

		sum = sum + result(opponent, me)
	}

	fmt.Printf("total score %d", sum)

}

func CreatingPairs(input string) game {
	games := make(game, len(strings.Split(input, "\n")))

	for k, g := range strings.Split(input, "\n") {
		v := strings.Split(g, " ")
		games[k] = append(games[k], v[0])
		games[k] = append(games[k], v[1])
	}
	return games
}

func result(opponent, me string) int {
	opChoice := opponent
	switch {
	case opChoice == "A":
		opChoice = "Rocks"
	case opChoice == "B":
		opChoice = "Paper"
	case opChoice == "C":
		opChoice = "Scissors"
	}

	myChoice := me
	score := 0
	switch {
	case myChoice == "Z":
		myChoice = "Scissors"
		score = 3
	case myChoice == "X":
		myChoice = "Rocks"
		score = 1
	case myChoice == "Y":
		myChoice = "Paper"
		score = 2
	}

	result := ""
	switch {
	case myChoice == opChoice:
		result = "draw"
		score = score + 3
	case myChoice == "Rocks" && opChoice == "Paper":
		result = "lost"
	case myChoice == "Paper" && opChoice == "Rocks":
		result = "win"
		score = score + 6
	case myChoice == "Paper" && opChoice == "Scissors":
		result = "lost"
	case myChoice == "Scissors" && opChoice == "Paper":
		result = "win"
		score = score + 6
	case myChoice == "Scissors" && opChoice == "Rocks":
		result = "lost"
	case myChoice == "Rocks" && opChoice == "Scissors":
		result = "win"
		score = score + 6
	}

	fmt.Printf("opponent choice: %s, my choice: %s = the result is: %s with a score of %d\n", opChoice, myChoice, result, score)
	return score
}
