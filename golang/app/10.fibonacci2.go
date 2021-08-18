package main

import "fmt"

func main()  {
	fmt.Println(fibonacci(10))
}

func fibonacci(n int) string{

	numb1 := 0
	numb2 := 1
	result := ""
	result += fmt.Sprintf("%d,%d", numb1, numb2)

	for i:=3; i <= n; i++ {
		output := numb1 + numb2
		result += fmt.Sprintf(",%d",output)

		numb1 = numb2
		numb2 = output
	}

	return result
}