package main

import "fmt"

var data = map[int]interface{}{
	1: "betul",
	2: true,
	5: true,
}

func main()  {
	input := []int{1,2,5}



	// fmt.Println(findTarget(input, 5))
	fmt.Println(findTargetAlt(input, 1))

}

func findTarget( input []int, target int) bool  {
	isFindTarget := false

	//dsds


	return isFindTarget

}

func findTargetAlt( targetData []int, target int) bool  {
	// isFindTarget := false
	 dataMap := make(map[int]bool)

	for _, value := range targetData {
		dataMap[value] = true
	}

	//dsds
	if value,ok := dataMap[target];ok{
		return true
	}


	return false

}