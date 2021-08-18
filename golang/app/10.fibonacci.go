package main

import "fmt"

func main()  {
	fmt.Println(fibonacci(10))
}

func fibonacci(n int) string{
	result := []int{}

	if n<0 {
		return ""
	}

	for i:=0; i<n; i++ {
		if i < 2 {
			result = append(result, i)
		} else {
			result = append(result, result[i-2] + result[i-1])
		}
	}

	stringResult := ""

	for i:=0; i< len(result); i++ {
		stringResult += fmt.Sprintf("%d ", result[i])
	}

	return stringResult
}