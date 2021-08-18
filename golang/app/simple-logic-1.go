/*
Buatlah program yang dapat menentukan grade dari suatu masukan angka, dengan kriteria sebagai berikut:
- Jika angka lebih besar atau sama dengan 90, maka output adalah huruf A
- Jika angka di antara 80 dan 89, maka output adalah huruf B
- Jika angka di antara 70 dan 79, maka output adalah huruf C
- Jika angka di antara 60 dan 69, maka output adalah huruf D
- Jika angka lebih kecil dari pada 59, maka output adalah huruf E
*/

package main

import "fmt"

func main()  {
	fmt.Println("SCORE", findScore(88))
}

func findScore(input int) string  {
	var score string
	if input>=90 {
		score = "A"
	}else if input>=80 && input<=90 {
		score = "B"
	}else if input>=70 && input<=80 {
		score = "C"
	}else if input>=60 && input<=70 {
		score = "D"
	}else if input<=59 {
		score = "E"
	}


	return score
}