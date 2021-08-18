package main

import "fmt"

func main()  {
	fmt.Println(fibonacci(10))
}

func fibonacci(n int) string{
	numb1 := 0
	numb2 := 1
	result := ""
	for i:=0; i < n; i++ {
		if i==0 {
			result += fmt.Sprintf("%d",i)
		} else if i==1{
			result += fmt.Sprintf(",%d",i)
		}else {
			output := numb1 + numb2
			result += fmt.Sprintf(",%d",output)
			numb1 = numb2
			numb2 = output
		}
	}
	return result
}