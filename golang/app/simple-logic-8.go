package main

import "fmt"

func main()  {
	input := []int{1,2,3}
	fmt.Println(findTotal(input))

}

func findTotal(input []int) int  {
	var total int
	for i:=0; i<len(input); i++ {
		if(input)
		total += input[i]
	}
	return total
}