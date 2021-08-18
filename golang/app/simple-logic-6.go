/*
Buatlah program rekursif. Anda bebas menentukan input dan output dari
fungsi tersebut, tetapi harus rekursif.
*/

package main

import "fmt"

func main()  {

	fmt.Println("recursive 1", recursive1(10))
	fmt.Println("recursive 2", recursive2(10))
	fmt.Println("recursive 3", recursive3(6))
}

// Penjumlahan
func recursive1(n int) int  {
	if(n==0) {
		return 0
	}

	return n + recursive1(n-1)
}

// Faktorial
func recursive2(n int) int  {
	if(n==0){
		return 1
	}
	return n * recursive2(n-1)
}

// Single Fibonacci
func recursive3(n int) int  {
	if(n<=1) {
		return n
	}
	return recursive3(n-1) + recursive3(n-2)
}