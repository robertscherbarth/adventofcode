package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	GetDistinctChar(string(content))

}

func GetDistinctChar(l string) {
	list := strings.Split(l, "")

	for i := 0; i < len(list); i++ {
		if i+14 > len(list) {
			break
		}

		if len(unique(list[i:i+14])) == 14 {
			fmt.Printf("length: %v; val: %v\n", len(unique(list[i:i+4])), list[i:i+4])
			index := strings.Index(l, strings.Join(list[i:i+14], ""))
			fmt.Println(index + 14)
			break
		}
	}
}

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
