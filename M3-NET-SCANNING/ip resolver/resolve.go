
package main

import (
	"net"
	"os"
	"fmt"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Resolve: %s IPv4|IPv6|HOSTNAME\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	fmt.Println("input:", name)
	parse := net.ParseIP(name)
	if parse == nil {
		// Resolve by domain name
		ips, err := net.LookupIP(name)
		if err != nil {
			fmt.Println("Resolution Error,", err.Error())
			os.Exit(1)
		}
		fmt.Println("Resolve by address:")
		for _,ip := range ips {
			fmt.Println("\t", ip.String())
		}
	} else {
		// Resolve by IP
		addrs, err := net.LookupAddr(name)
		if err != nil {
			fmt.Println("Resolution Error,", err.Error())
			os.Exit(1)
		}
		fmt.Println("Resolve by IP:")
		for _,addr := range addrs {
			fmt.Println("\t", addr)
		}
	}
}