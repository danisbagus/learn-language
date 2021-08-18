package main

import (
	"encoding/json"
	"fmt"
)

type Input struct {
	Id int
	Name string
	Age int
}

func main()  {
	inputJson:= `
	[
		{ "id": 1, "name": "Udin", "age": 12 },
		{ "id": 2, "name": "Reane", "age": 51 },
		{ "id": 3, "name": "Budi", "age": 34 },
		{ "id": 4, "name": "Agus", "age": 16 },
		{ "id": 5, "name": "Sari", "age": 19 },
		{ "id": 6, "name": "Ririn", "age": 20 },
		{ "id": 7, "name": "Dessy", "age": 23 }
	]
	`
	var input []Input

	json.Unmarshal([]byte(inputJson), &input)
	// fmt.Printf("Input : %+v", input)

	fmt.Printf("filtered data : %+v", findDataByAge(input, 21))
}

func findDataByAge(input []Input, age int) []Input  {
	var filteredData []Input

	for i:=0; i<len(input); i++ {
		if(input[i].Age < age) {
			filteredData = append(filteredData, input[i] )
		}
	}

	return filteredData
}