

package main

import (
	"fmt"
	"net"
)


func main() {
	// https://golang.org/pkg/net/#Dial
	_, err := net.Dial("tcp", "scanme.nmap.org:23")
	if err != nil {
		fmt.Println("Port 23 is closed/filtered.")
		fmt.Println(err)
	} else {
		fmt.Println("Port 23 is open.")
	}
}