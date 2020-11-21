
package main

import (
	"fmt"
	// ..

)

func main() {
	// Line comments
	/*
		General comments
	*/
	// GPA := 18
	// fmt.Print("Hi, Gophers!\n")
	// fmt.Println("A", "B", "C", "Z")
	// fmt.Println("Hi Gophers!")
	// fmt.Println("My name is " + "Saeed!")
	// fmt.Printf("How old are you? Im %d", 19)
	// fmt.Printf("Your GPA is %d.\n", GPA)
	desc := fmt.Sprintf("Totally: %d\nFrequency: %s\nA float: %f\nCHR: %c", 10, "ASKED", 2.2, 'c')
	fmt.Println(desc)
}