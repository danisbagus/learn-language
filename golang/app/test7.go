package main

import "fmt"

func main()  {
	input := []int{3,5,2,-4,8,11}

	fmt.Println("result: ", findPairNumbers(input, 13))
}

func findPairNumbers(input []int, target int) [][]int  {

	var result [][]int

	//slice to map
	var mapInput = make(map[int]bool)
	for i:=0; i<len(input); i++{
		mapInput[input[i]] = true
	}

	for i:=0; i<len(input); i++{
		sisa := target - input[i]
		if _, isFound := mapInput[sisa]; isFound{
			result = append(result, []int{input[i],sisa})
		}
	}

	return result
}