
package main

import (
	"log"
	"strconv"
	"net"
	"sync"
)

var IP = "129.42."

func main() {
	var WG sync.WaitGroup
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			WG.Add(1)
			go func(ip string) {
				defer WG.Done()
				lookup, err := net.LookupAddr(ip)
				if err == nil {
					log.Println(ip)
					for _,addr := range lookup {
						log.Println("\t", addr)
					}
				}
			}(IP + strconv.Itoa(i) + "." + strconv.Itoa(j))
		}
	}
	WG.Wait()
}