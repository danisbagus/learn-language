package main

import "fmt"

func main()  {
	result:= factorial(5)
	fmt.Printf("Factorial: %d\n", result)
}

func factorial(numb int) int  {
	if(numb == 0 ) {
		return 1
	}
	return numb * factorial(numb - 1)
}