package main

import "fmt"

func main()  {
    diameter:= 20
    radius:= 10.0
    const phi = 3.14
    var isUsingRadius bool = true

    var result float64

    if(isUsingRadius){
      result = calculate2(radius)
    } else {
      result = calculate1(diameter)
    }

    fmt.Printf("Circle area is: %f\n", result)

}

func calculate1(diameter int) float64 {
  radius:= float64(diameter/2)
  result:= 3.14 * radius * radius
  return result
}

func calculate2(radius float64) float64  {
  result:= 3.14 * radius * radius
  return result
}
