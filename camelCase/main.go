package main

import (
	"fmt"
	"strings"
)

func countWords1(stringInput string) int {
	answer := 1
	for _, ch := range stringInput {
		min, max := 'A', 'Z'
		if ch >= min && ch <= max {
			answer++
		}
	}
	return answer
}

func countWords2(stringInput string) int {
	answer := 1
	for _, ch := range stringInput {
		charStr := string(ch)
		if strings.ToUpper(charStr) == charStr {
			answer++
		}
	}
	return answer
}

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Println(input)

	answer := countWords2(input)
	fmt.Println(answer)
}
