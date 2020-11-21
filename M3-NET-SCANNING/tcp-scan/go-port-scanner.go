

package main

import (
	"strconv"
	"log"
	"net"
	"sync"
)


var (
	SERVER = "scanme.nmap.org"
	RANGE = 1024
)

func main() {
	var wg sync.WaitGroup
	// https://golang.org/pkg/net/#Dial
	for port := 1; port <= RANGE; port++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			req, err := net.Dial("tcp", SERVER + ":" + strconv.Itoa(i))
			if err == nil {
				log.Printf("Port %d is open.\n", i)
				req.Close()
			}
		}(port)
	}
	wg.Wait()
}