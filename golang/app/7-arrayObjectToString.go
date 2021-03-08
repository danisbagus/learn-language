package main

import "fmt"

func main()  {

	var arrayPoint = []map[string]float64{
		{ "x": 34.79600143432617, "y": 700.9229736328125},
		{ "x": 71.913002014160167, "y": 737.7329711914062},
		{ "x": 426.864990234375, "y": 737.7670288085938},
		{ "x": 425.6109924316406, "y": 529.6099853515625},
		{ "x": 387.1189880371094, "y": 528.7100219726562},
		{ "x": 386.4949951171875, "y": 448.89599609375},
	}

	arrayLength := len(arrayPoint);

	var stringPoint string;

	for i:=0; i < arrayLength; i++ {
		stringPoint += fmt.Sprintf("%f %f ", arrayPoint[i]["x"], arrayPoint[i]["y"])
	}

	fmt.Println(stringPoint)
}