package main

import (
	"strconv"
	"log"
	"net"
	"time"
)

var (
	SERVER = "scanme.nmap.org"
	RANGE = 1024
	TIMEOUT = time.Second * 20
)

func main() {
	active_threads := 0
	finished_tasks := make(chan bool)
	for port := 1; port <= RANGE; port++ {
		go func(i int){
			_, err := net.DialTimeout("tcp", SERVER + ":" + strconv.Itoa(i), TIMEOUT)
			if err == nil {
				log.Printf("Port %d is open.\n", i)
			}
			finished_tasks <- true
		}(port)
		active_threads++
	}
	for active_threads > 0 {
		<- finished_tasks
		active_threads--
	}
}
