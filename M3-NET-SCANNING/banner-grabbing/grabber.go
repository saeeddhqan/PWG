

package main

import (
	"strconv"
	"log"
	"net"
	"sync"
	"time"
)


var (
	SERVER = "127.0.0.1"
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
			if err != nil {
				return
			}
			defer req.Close()
			log.Printf("Port %d is open.\n", i)
			buf := make([]byte, 4094)
			req.SetReadDeadline(time.Now().Add(time.Second * 8))
			read, err := req.Read(buf)
			if err != nil {
				return
			}
			log.Printf("Banner(%d): %s", i, buf[:read])

		}(port)
	}
	wg.Wait()
}
