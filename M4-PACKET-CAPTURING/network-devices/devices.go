

package main

import (
	"fmt"
	"log"
	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}

	for _, device := range devices {
		fmt.Println(device.Name, ":", device.Description)
		fmt.Println()
		for _, address := range device.Addresses {
			fmt.Println("\t\tIP:", address.IP)
			fmt.Println("\t\tNet Mask:", address.Netmask)
			fmt.Println("\t\tBroad Address:", address.Broadaddr)
			fmt.Println("\t\tP2P:", address.P2P)
			fmt.Println("-----------")
		}
	}
}
