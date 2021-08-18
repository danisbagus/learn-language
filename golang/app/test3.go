package main

import "fmt"

func main()  {
	fmt.Println("result",findFPB(16, 40))
	fmt.Println("result",findKPK(4, 6))
}

func findFPB(a int, b int) int  {
	n := b
	for n > 0 && !(a%n == 0 && b%n == 0){
		n--
	}
	return n
}

func findKPK(a int, b int) int  {
	n := b
	for n > 0 && n % a != 0 || n % b != 0 {
		n++
	}
	return n;
}



