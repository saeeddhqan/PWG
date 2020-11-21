
package main

import (
	"fmt"
)


type Person struct {
	Name 	string
	Age		int
}


type Student struct {
	person 			Person
	fieldOfStudy 	string
	department		string
}


func (p *Person) sayYourName() {
	fmt.Println(p.Name)
}

func (p *Person) sayYourAge() {
	fmt.Println(p.Age)
}

func (p *Person) ChangeYourAge(age int) {
	p.Age = age
}


func main() {

	John := Person{Name: "John", Age : 38}
	// var John Person = Person{Name: "John", Age : 38}
	// John.Age = 40
	// fmt.Println(John.Age, John.Name)
	// S1 := Student{person: John, fieldOfStudy: "NLP", department: "CS"}
	// fmt.Println(S1.person.Name)
	// fmt.Println(S1.person.Age)
	// fmt.Println(S1.fieldOfStudy)
	// fmt.Println(S1.department)
	//////////////////////////////////
	// John.sayYourName()
	// John.ChangeYourAge(40)
	// John.sayYourAge()

}