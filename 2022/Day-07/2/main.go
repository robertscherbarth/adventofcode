package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type Stack[T any] struct {
	sync.Mutex
	stack []T
}

func New[T any]() *Stack[T] {
	return &(Stack[T]{stack: []T{}})
}

func (s *Stack[T]) Len() int {
	return len((*s).stack)
}

func (s *Stack[T]) Push(item T) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	(*s).stack = append((*s).stack, item)
}

func (s *Stack[T]) Pop() (item T) {
	if len(s.stack) == 0 {
		return
	}
	s.Mutex.Lock()

	item = (*s).stack[len((*s).stack)-1]
	(*s).stack = (*s).stack[:len((*s).stack)-1]

	s.Mutex.Unlock()

	return item
}

func (s *Stack[T]) Peak() (item T) {
	if len(s.stack) == 0 {
		return
	}
	s.Mutex.Lock()
	item = (*s).stack[len((*s).stack)-1]
	s.Mutex.Unlock()
	return item
}

func (s *Stack[T]) GetStack() (stack []T) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return (*s).stack
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	sizes := make(map[string]int, 0)
	stack := New[string]()
	regex := "([0-9]+)"
	matcher := regexp.MustCompile(regex)

	for _, line := range strings.Split(string(content), "\n") {
		switch {
		case strings.Contains(line, "$ cd .."):
			stack.Pop()
		case strings.Contains(line, "$ cd"):
			dir := strings.Split(line, " ")[2]
			stack.Push(fmt.Sprintf("%s%s", stack.Peak(), dir))
			continue
		case matcher.MatchString(line):
			value, _ := strconv.Atoi(strings.Split(line, " ")[0])
			for _, dir := range stack.GetStack() {
				sizes[dir] = sizes[dir] + value
			}
		}
	}

	freeSpace := 70000000 - sizes["/"]
	list := make([]int, 0)
	for _, v := range sizes {
		if v+freeSpace >= 30000000 {
			list = append(list, v)
		}
	}

	min := list[0]
	for _, v := range list {
		if min > v {
			min = v
		}
	}

	fmt.Printf("minimal value %d\n", min)
}
