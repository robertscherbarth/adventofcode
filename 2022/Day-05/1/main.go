package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// stacks
type stacks struct {
	Queue map[int][]string
}

func (s *stacks) Push(column int, elem string) {
	s.Queue[column] = append(s.Queue[column], elem)
}

func (s *stacks) Pop(column int) string {
	elem := s.Queue[column][len(s.Queue[column])-1]
	s.Queue[column] = s.Queue[column][:len(s.Queue[column])-1]
	return elem
}

func (s *stacks) LastItem(column int) string {
	return s.Queue[column][len(s.Queue[column])-1]
}

type command struct {
	move int
	from int
	to   int
}

func main() {
	stackQueue := stacks{Queue: map[int][]string{}}
	stackQueue.Queue[1] = []string{"N", "D", "M", "Q", "B", "P", "Z"}
	stackQueue.Queue[2] = []string{"C", "L", "Z", "Q", "M", "D", "H", "V"}
	stackQueue.Queue[3] = []string{"Q", "H", "R", "D", "V", "F", "Z", "G"}
	stackQueue.Queue[4] = []string{"H", "G", "D", "F", "N"}
	stackQueue.Queue[5] = []string{"N", "F", "Q"}
	stackQueue.Queue[6] = []string{"D", "Q", "V", "Z", "F", "B", "T"}
	stackQueue.Queue[7] = []string{"Q", "M", "T", "Z", "D", "V", "S", "H"}
	stackQueue.Queue[8] = []string{"M", "G", "F", "P", "N", "Q"}
	stackQueue.Queue[9] = []string{"B", "W", "R", "M"}

	content, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	commands := InitCommands(string(content))

	for _, v := range commands {
		for i := 0; i < v.move; i++ {
			elem := stackQueue.Pop(v.from)
			stackQueue.Push(v.to, elem)
		}
	}

	for i := 1; i <= len(stackQueue.Queue); i++ {
		fmt.Printf("%s", stackQueue.LastItem(i))

	}
}

func InitCommands(content string) []command {
	list := strings.Split(content, "\n")
	commands := make([]command, 0)

	for _, v := range list {
		fields := strings.Fields(v)
		move, _ := strconv.Atoi(fields[1])
		from, _ := strconv.Atoi(fields[3])
		to, _ := strconv.Atoi(fields[5])
		commands = append(commands, command{
			move: move,
			from: from,
			to:   to,
		})
	}

	return commands
}
