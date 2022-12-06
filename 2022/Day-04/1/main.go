package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ID number, and each Elf is assigned a range of section IDs.

// sections
type sections [][]ID

type ID struct {
	LowerLimit int
	UpperLimit int
}

func (i *ID) InRange(upper, lower int) bool {
	var inrange bool

	if upper >= i.UpperLimit && lower <= i.LowerLimit {
		inrange = true
	}
	if i.UpperLimit >= upper && i.LowerLimit <= lower {
		inrange = true
	}

	return inrange
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	s := CreateSectionsPairs(string(content))
	fmt.Println(s)

	overlaps := GetOverlaps(s)
	fmt.Println(overlaps)

}

func CreateSectionsPairs(content string) sections {
	list := strings.Split(content, "\n")
	ship := make(sections, len(list))
	var x int
	for _, v := range list {
		for _, l := range strings.Split(v, ",") {
			s := strings.Split(l, "-")
			lowerLimit, _ := strconv.Atoi(s[0])
			upperLimit, _ := strconv.Atoi(s[1])
			id := ID{
				LowerLimit: lowerLimit,
				UpperLimit: upperLimit,
			}
			ship[x] = append(ship[x], id)
		}
		x = x + 1
	}

	return ship
}

func GetOverlaps(sections sections) int {
	var counter int
	for _, v := range sections {
		if v[0].InRange(v[1].UpperLimit, v[1].LowerLimit) {
			counter = counter + 1
		}
	}
	return counter
}
