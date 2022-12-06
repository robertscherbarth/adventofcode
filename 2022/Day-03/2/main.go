package main

import (
	"fmt"
	"os"
	"strings"
)

type rucksacks [][]string

func main() {

	content, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	ruck := CreateRucksacks(string(content))
	fmt.Println(ruck)
	prio := GetPrio(ruck)

	var sum int
	for _, v := range prio {
		sum = sum + prioritizeType(v)
		fmt.Printf("char: %s prio %d\n", string(v), prioritizeType(v))
	}

	fmt.Printf("sum over prio %d\n", sum)

}

func CreateRucksacks(content string) rucksacks {
	list := strings.Split(content, "\n")
	ruck := make(rucksacks, len(list)/3)
	var x, y int
	for i := 0; i < len(list); i++ {
		if x > 2 {
			y = y + 1
			x = 0

		}

		ruck[y] = append(ruck[y], list[i])
		x = x + 1
	}

	return ruck
}

func GetPrio(ruck rucksacks) []int32 {
	res := make([]int32, 0)

	for _, r := range ruck {
		for _, c := range r[0] {

			if strings.Contains(r[1], string(c)) && strings.Contains(r[2], string(c)) {
				res = append(res, c)
				break
			}
		}
	}

	return res
}

func prioritizeType(itemType int32) int {
	val := itemType - 38
	if itemType >= 97 {
		val = itemType - 96
	}
	return int(val)
}
