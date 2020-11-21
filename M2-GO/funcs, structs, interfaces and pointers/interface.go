
package main

import (
	"fmt"
)

type Student interface {
	sayYourName()
	sayHello()
}

type Person struct {
	Name 	string
	Age		int
}

func (p Person) sayYourName() {
	fmt.Println("I'm", p.Name)
}

func (p Person) sayHello() {
	fmt.Println("Hello!")
}
func main() {

	var student Student
	student = Person{"Foo", 10}
	student.sayYourName()
	student.sayHello()

}