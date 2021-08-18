/*
Buatlah program fungsi untuk menentukan suatu bilangan merupakan bilangan ganjil atau genap.
*/

package main

import "fmt"

func main()  {
	fmt.Println(checkBilangan(19))
	fmt.Println(checkBilangan(20))
}

func checkBilangan(input int) string  {
	if(input%2 == 0) {
		return "Genap"
	}else{
		return "Ganjil"
	}
}