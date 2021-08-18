package main

import (
	"fmt"
	"strings"
)

func main()  {
	input := "ayo belajar golang"
	fmt.Println(reverse(input))
}

func reverse(input string)  string{
	arrayInput := strings.Split(input, " ")
	reverseInput := []string{}

	for i:=0; i<len(arrayInput); i++{
		reverseInput = append(reverseInput, arrayInput[len(arrayInput)-1-i])
	}

	result := strings.Join(reverseInput, " ")
	return result
}