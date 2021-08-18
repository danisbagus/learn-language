package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main()  {
	Person1 := Person{"Julio", 20}

	fmt.Println("Name: ", Person1.Name)
	// fmt.Println("Age: ", Person1.Age)
	setAge(&Person1, 22)
	fmt.Println("Age: ", getAge(Person1))

}

func getAge(person Person) int  {
	return person.Age
}

func setAge(person *Person, age int)  {
	person.Age = age
}

// func setAge(person Person, age int)  {
// 	person.Age = age
// }