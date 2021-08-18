package main

import "fmt"

func main()  {
	fmt.Println(fibonacci(10))
	fmt.Println(recursive(6))

}

func fibonacci(n int) []int {
	result := []int{}
	for i := 0; i < 10; i++ {
		result = append(result, recursive(i))
	}
	return result
}

func recursive(n int) int{
	if(n<=1) {
		return n
	}

	return recursive(n-1) + recursive(n-2)
}

