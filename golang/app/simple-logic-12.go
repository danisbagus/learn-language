package main

import "fmt"

func main()  {
	input := []int{15,1,2}

	fmt.Println(findNearestFibonacci(input))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findNearestFibonacci(input []int) int {
	total := 0
	for i:=0; i<len(input); i++ {
		total += input[i]
	}

	isFindNearest := false
	i := 0
	result := 0
	for !isFindNearest {
		selisih := recursive(i) - total
		selisih2 := recursive(i + 1) - total

		if( Abs(selisih) < Abs(selisih2)) {
			result = selisih
			isFindNearest = true
		}
		i++
	}
	return result
}

func recursive(n int) int{
	if(n<=1) {
		return n
	}
	return recursive(n-1) + recursive(n-2)
}

