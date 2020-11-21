
package main

import (
	"net"
	"log"
	"fmt"
)


func main() {
	req, err := net.Dial("tcp", "localhost:2020")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(req, "Hi this is client!")
	buf := make([]byte, 4096)
	read, err := req.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Received data: %s\n", buf[:read])
}