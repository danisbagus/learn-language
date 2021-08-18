package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("is palindrom: ", checkPalindrom("kasar"))
	fmt.Println("is palindrom: ", checkPalindrom("rasar"))
}

func checkPalindrom(input string) bool  {
	inputData := strings.ToLower(input)

	for i:=0; i < len(inputData); i++ {
		if(inputData[i] != inputData[len(inputData)-1-i]){
			return false
		}
	}
	return true
}