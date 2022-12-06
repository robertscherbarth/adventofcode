package main

import (
	"fmt"
	"strconv"
	"strings"
)

type elvesMap map[int]int

func main() {

	elves := make(elvesMap, 0)

	input := `...`

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

	fmt.Printf("number of elves %d\n", len(elves))

	topThree := make(elvesMap, 0)
	for {
		if len(topThree) == 3 {
			break
		}
		elv, cal := mostCalories(elves)
		topThree[elv] = cal
		delete(elves, elv)

		fmt.Printf("the elv %d with the most calories %d\n", elv, cal)
	}

	sum := 0
	for _, v := range topThree {
		sum += v
	}

	fmt.Printf("sum of the top three %d\n", sum)

}

func mostCalories(elves elvesMap) (int, int) {
	calories := 0
	elve := 0
	for k, v := range elves {
		if calories < v {
			calories = v
			elve = k
		}
	}
	return elve, calories
}
