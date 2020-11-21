
package main

import (
	"fmt"
)




func main() {
	var num = int(1024)
	var point *int = &num
	fmt.Println(*point)
	*point = 512
	fmt.Println(num)
}