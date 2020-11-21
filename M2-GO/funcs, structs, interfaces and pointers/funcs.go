
package main


import (
	"fmt"
)


// func add(x int, y int) (int) {
// 	return x + y
// }

// func actionShort(x, y int) (int, int, int, int) {
// 	return x + y, x * y, x / y, x - y
// }


// func getAction(op string) (func(int, int) (int)) {
// 	if op == "+" {
// 		return func(a, b int) (int) {return a + b}
// 	} else if op == "*" {
// 		return func(a, b int) (int) {return a * b}
// 	} else if op == "/" {
// 		return func(a, b int) (int) {return a / b}
// 	} else {
// 		return func(a, b int) (int) {return a - b}
// 	}
// }

// func action(x int, y int, ac func(int, int) (int)) (int) {
// 	return ac(x, y)
// }

// func variadicParam(seq ...string) {
// 	fmt.Println(seq)
// }

// func sumVariadicParam(seq ...int) (int) {
// 	sum := 0
// 	for _,v := range seq {
// 		sum += v
// 	}
// 	return sum
// }

func main() {
	// fmt.Println(actionShort(5, 10))
	// a, b, c, d := actionShort(5, 10)
	// fmt.Println(a, b, c, d)
	// addFunc := func(a, b int) (int) {return a + b}
	// var addFunc func(int, int) (int) = func(a, b int) (int) {return a + b}
	// var mulFunc func(int, int) (int) = func(a, b int) (int) {return a * b}
	// var divFunc func(int, int) (int) = func(a, b int) (int) {return a / b}
	// var minFunc func(int, int) (int) = func(a, b int) (int) {return a - b}
	// fmt.Println(addFunc(5, 10))
	// fmt.Println(action(5, 10, getAction("+")))
	// fmt.Println(action(5, 10, getAction("*")))
	// fmt.Println(action(5, 10, getAction("/")))
	// fmt.Println(action(5, 10, getAction("-")))
	// fmt.Println(action(5, 10, mulFunc))
	// fmt.Println(action(5, 10, divFunc))
	// fmt.Println(action(5, 10, minFunc))

	// variadicParam("a", "b", "c", "x", "x", "x", "X")
	// fmt.Println(sumVariadicParam(1, 10, 100, 1, 1, 1))
}