/*
Buatlah program untuk menentukan nilai maksimum, minimum, nilai tengah, dan rata-rata dari suatu deretan angka.
Contoh:
-6, -4, 0, 1, 2, 2, 3, 10, 12, 44, 52, 72, 98

Max:     98

Min:     -6

Median:   3

Average: 22

*/

package main

import "fmt"

func  main()  {
	input := []int{-6, -4, 0, 1, 2, 2, 3, 10, 12, 44, 52, 72, 98}
	fmt.Println("max:", findMax(input))
	fmt.Println("min:", findMin(input))
	fmt.Println("avg:", findAverage(input))
}

func findMax(input []int) int  {
	maxValue := 0
	for i:=0; i<len(input); i++{
		if(input[i] > maxValue) {
			maxValue = input[i]
		}
	}
	return maxValue
}

func findMin(input []int) int  {
	minValue := 0
	for i:=0; i<len(input); i++{
		if(input[i] < minValue) {
			minValue = input[i]
		}
	}
	return minValue
}

func findAverage(input []int) int  {
	total :=0
	avg := 0

	for i:=0; i<len(input); i++{
		total += input[i]
	}

	if total!=0 {
		avg = total / len(input)
	}

	return avg
}