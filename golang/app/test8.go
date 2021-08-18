package main

import "fmt"

func main()  {
	fmt.Println(findFibonacci(10))
	fmt.Println(findFibonacci2(10))
}

func findFibonacci(n int) []int {
	result := []int{}

	for i:=0; i<n; i++{
		if(i<2) {
			result = append(result,i)
		}else{
			result = append(result, result[i-2]+result[i-1])
		}
	}

	return result
}

func findFibonacci2(n int) string  {
	numb1 := 0
	numb2 :=1
	result := fmt.Sprintf("%d,%d", numb1, numb2)

	for i:=2; i<n; i++ {
		jumlah := numb1 + numb2
		numb1 = numb2
		numb2 = jumlah
		result += fmt.Sprintf(",%d", jumlah)
	}
	return result
}