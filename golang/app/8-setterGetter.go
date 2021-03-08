package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main()  {
	var person1 = Person{"Hera", 20}
	person1.getName();
	person1.setName("Hero")
	person1.getName();
}

func (p *Person) getName()  {
	fmt.Println("Get name: ", p.Name)
}

func (p* Person) setName(name string) {
	p.Name = name
}

