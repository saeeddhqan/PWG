
package main

import (
	"fmt"
)


func main() {
	// x, y := 1, 10

	// if !(x == 1) {
	// 	fmt.Println("Condition is true")
	// } else {
	// 	fmt.Println("Condition is false")
	// }

	// if x != 1 {
	// 	fmt.Println("x is not equal to one.")
	// } else {
	// 	fmt.Println("x is equal to one.")
	// }

	// if x == 1 && y == 10 {
	// 	fmt.Println("Condition is true")
	// } else {
	// 	fmt.Println("Condition is false")
	// }

	// num := 100010

	// if num < 1024 {
	// 	fmt.Println("Num is less than 1024.")
	// } else if num < 10000 {
	// 	fmt.Println("Num is less than 10000.")
	// } else if num > 10000 && num < 20000 {
	// 	fmt.Println("Num is betwixt 10000 and 20000.")
	// } else {
	// 	fmt.Println("There is no condition to run.")
	// }

	// if num := 1024; num < 1024 {
	// 	fmt.Println("Num is less than 1024.")
	// } else if num < 10000 {
	// 	fmt.Println("Num is less than 10000.")
	// } else if num > 10000 && num < 20000 {
	// 	fmt.Println("Num is betwixt 10000 and 20000.")
	// } else {
	// 	fmt.Println("There is no condition to run.")
	// }

	// fmt.Println(num)

	x := 11

	// switch x {
	// 	case 1:
	// 		fmt.Println("One")
	// 	case 2:
	// 		fmt.Println("Two")
	// 		fallthrough
	// 	case 3:
	// 		fmt.Println("Three")
	// 	case 4:
	// 		fmt.Println("Four")
	// 	// case "Five":
	// 	// 	fmt.Println("Five string")
	// 	default:
	// 		fmt.Println("There is no case for this number.")
	// }

	switch {
		case x == 1:
			fmt.Println("One")
		case x == 1 || x == 2:
			fmt.Println("One or Two")
		case x > 10 && x < 20:
			fmt.Println("X is betwixt 10 and 20")
		default:
			fmt.Println("There is no case for this number.")
	}

}