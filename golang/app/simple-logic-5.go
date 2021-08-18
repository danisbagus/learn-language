/*
Buatlah program yang memiliki dua input berupa tahun. Output dari fungsi tersebut
adalah list dari tahun-tahun kabisat diantara dua input tahun tersebut.
*/

package main

import "fmt"

func main()  {
	fmt.Println(findKabisat(2000, 2008))
}

func findKabisat(n1 int, n2 int) []int {

	result := []int{}

	for i:=n1; i<=n2; i++ {
		if(i%4 == 0) {
			result = append(result, i)
		}
	}

	return result
}