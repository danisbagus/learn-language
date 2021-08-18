/*
	Buatlah program untuk menentukan suatu kata atau kalimat merupakan palindrom atau bukan.
	Contoh:


	Radar                 # --> palindrom
	Malam                 # --> palindrom
	Kasur ini rusak       # --> palindrom
	Ibu Ratna antar ubi   # --> palindrom

	Malas                 # --> bukan palindrom
	Makan nasi goreng     # --> bukan palindrom
	Balonku ada lima      # --> bukan palindrom
*/

package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(checkPalindrom("Radar"))
	fmt.Println(checkPalindrom("Kasur ini rusak"))
	fmt.Println(checkPalindrom("Malas"))
	fmt.Println(checkPalindrom("Balonku ada lima"))

	fmt.Println(checkPalindrom2("Radar"))
	fmt.Println(checkPalindrom2("Kasur ini rusak"))
	fmt.Println(checkPalindrom2("Malas"))
	fmt.Println(checkPalindrom2("Balonku ada lima"))
}

func checkPalindrom(input string) bool  {
	inputData := strings.ToLower(input)

	for i:=0; i<len(inputData); i++ {
		if inputData[i] != inputData[len(inputData)-i-1] {
			return false
		}
	}
	return true
}

func checkPalindrom2(input string) bool  {
	var reverseInput string

	inputData := strings.ToLower(input)

	for i:=0; i<len(inputData); i++ {
		reverseInput += fmt.Sprintf("%s", string(inputData[len(inputData)-i-1]))
	}

	if(reverseInput != inputData) {
		return false
	}

	return true
}