package main

import "fmt"

func main()  {
	input := []int{15,1,2}

	fmt.Println(nearFibonacci(input))
	// fmt.Println("check", fibonacci(9))
}

func nearFibonacci(input []int) int {
	var total int = 0
	var isNear bool = false
	var selisih int = 0
	var i int = 0

	for i:=0; i<len(input); i++ {
		total += input[i]
	}

	for !isNear {
		selisih1 := Abs(total - fibonacci(i))
		selisih2 := Abs(total -  fibonacci(i+1))

		if selisih1 < selisih2 {
			isNear = true
		}
		selisih = selisih1
		i++
	}

	return selisih
}

func fibonacci(n int) int  {
	if(n<2){
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func Abs(input int) int  {
	if(input < 0){
		return -input
	}
	return input
}

// 0,1,1,2,3,5,8,13,21,34