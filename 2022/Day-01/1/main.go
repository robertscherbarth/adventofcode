package main

import (
	"fmt"
	"strconv"
	"strings"
)

type elves map[int]int

func main() {

	elves := make(elves, 0)

	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	fmt.Printf("size of input %d\n", len(strings.Fields(input)))
	c := 1
	calories := 0

	for k, v := range strings.Split(input, "\n") {
		i, _ := strconv.Atoi(v)
		calories = calories + i
		elves[c] = calories

		if v == "" {
			fmt.Printf("number of elves %d with total calories %d\n", c, calories)
			c++
			calories = 0
		}

		if len(strings.Split(input, "\n"))-1 == k {
			elves[c] = calories
			fmt.Printf("number of elves %d with total calories %d\n", c, calories)
		}
	}

	fmt.Printf("number of elves %d", len(elves))

	mostCalories := 0
	for _, v := range elves {
		if mostCalories < v {
			mostCalories = v
		}
	}

	fmt.Printf("the most calories %d", mostCalories)
}
