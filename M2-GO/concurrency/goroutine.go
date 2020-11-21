
package main

import (
	"log"
	"time"
	"math"
)

func count(add string) {
	for i := 0; i < 10; i++ {
		log.Println(i, add)
		time.Sleep(time.Millisecond * 500)
	}
}

func pow(x, y float64, resp chan float64) {
	resp <- math.Pow(x, y)
}

func main() {
	// go count("-")
	// count("+")
	// time.Sleep(time.Second * 2)
	// ch <- data
	// data <- ch

	resp := make(chan float64)
	go pow(2, 2, resp)
	go pow(2, 10, resp)
	second, first := <-resp, <-resp
	log.Println(first)
	log.Println(second)
}