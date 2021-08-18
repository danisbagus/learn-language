package main

import "fmt"

func main()  {
	// input := [][]int{{1,2},{3,4}}
	// input := [][]int{}
	// testing := append(input, []int{1,2,5})
	// testing = append(testing, []int{11,12,15})

	input := []int{3,5,2,-4,8,11}

	fmt.Println("result: ", findPairNumbs(input, 13))
}

func findPairNumbs( numbs []int, target int) [][]int {
	result := [][]int{}

	// Convert array to maps
	mapInput := make(map[int]bool)
	for i:=0; i <len(numbs); i++ {
		// mapInput[i] = numbs[i]
		mapInput[numbs[i]] = true
	}

	for i:=0; i < len(numbs); i++ {
		selisih := target - numbs[i]
		if _,found := mapInput[selisih]; found {
			result = append(result, []int{numbs[i], selisih})
		}
	}
	return result
}