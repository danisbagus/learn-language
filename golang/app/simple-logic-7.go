/*
Buatlah program untuk membalikkan "semua kata-kata" (bukan huruf) yang terdapat dalam suatu kalimat.

Contoh:

# Input
makan nasi goreng

# Output
goreng nasi makan
*/

package main

import (
	"fmt"
	"strings"
)

func main()  {
	input := "makan nasi goreng"
	fmt.Println(reverse(input))
}

func reverse(input string) string {
	arrayInput := strings.Split(input, " ")
	var reverseInput string
	for i:=0; i<len(arrayInput); i++ {
		reverseInput += fmt.Sprintf("%s ", arrayInput[len(arrayInput) - i - 1])
	}
	return reverseInput
}