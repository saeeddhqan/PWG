

package main

import (
	"fmt"
	"log"
	"net"
)


var (
	SERVER = "127.0.0.1"
	RANGE = 65535
)

func main() {
	// https://golang.org/pkg/net/#Dial
	for port := 1; port <= RANGE; port++ {
		_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", SERVER, port))
		if err == nil {
			log.Printf("Port %d is open.\n", port)
		}
	}
}