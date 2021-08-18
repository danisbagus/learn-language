package main

import (
	"fmt"
)

func main()  {
	// fmt.Println("10",checkPrima(10))
	fmt.Println(findBilanganPrima(10))

}

func findBilanganPrima(n int) string {

	result := []int{}
	i := 0
	for len(result) < n {
		if(checkPrima(i)){
			result = append(result, i)
		}
		i++
	}
	// fmt.Printf(strings.Join(result,','))
	var stringResult string

	for i:=0; i<len(result); i++{
		stringResult += fmt.Sprintf("%d ", result[i])
	}

	return stringResult
}

func checkPrima(n int) bool  {
	if(n < 2) {
		return false
	}
	for i:=2; i <= n/2; i++ {
		if(n%i == 0){
			return false
		}
	}
	return true
}

//2,3,5,7,11,13,17,19