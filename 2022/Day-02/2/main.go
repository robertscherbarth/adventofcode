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

func result(opponent, exp string) int {
	opChoice := opponent
	switch {
	case opChoice == "A":
		opChoice = "Rocks"
	case opChoice == "B":
		opChoice = "Paper"
	case opChoice == "C":
		opChoice = "Scissors"
	}

	myChoice := ""
	score := 0

	switch {
	case exp == "Y" && opChoice == "Rocks":
		myChoice = "Rocks"
	case exp == "Y" && opChoice == "Paper":
		myChoice = "Paper"
	case exp == "Y" && opChoice == "Scissors":
		myChoice = "Scissors"
	case exp == "X" && opChoice == "Rocks":
		myChoice = "Scissors"
	case exp == "X" && opChoice == "Scissors":
		myChoice = "Paper"
	case exp == "X" && opChoice == "Paper":
		myChoice = "Rocks"
	case exp == "Z" && opChoice == "Paper":
		myChoice = "Scissors"
	case exp == "Z" && opChoice == "Rocks":
		myChoice = "Paper"
	case exp == "Z" && opChoice == "Scissors":
		myChoice = "Rocks"
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

	switch {
	case myChoice == "Scissors":
		score = score + 3
	case myChoice == "Rocks":
		score = score + 1
	case myChoice == "Paper":
		score = score + 2
	}

	fmt.Printf("opponent choice: %s, my choice: %s = the result is: %s with a score of %d\n", opChoice, myChoice, result, score)
	return score
}
