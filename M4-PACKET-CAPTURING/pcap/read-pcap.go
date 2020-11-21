
package main

import (
	"fmt"
	"log"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"flag"
)

var filename = flag.String("f", "", "PCAP file name")

func main() {
	flag.Parse()
	if *filename == "" {
		log.Fatalln("The file name is not specified!")
	}
	handler, err := pcap.OpenOffline(*filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer handler.Close()

	for packet := range gopacket.NewPacketSource(handler, handler.LinkType()).Packets() {
		// fmt.Println(packet) // packet.String()
		// fmt.Println(packet.Dump())
		fmt.Println(packet)
		fmt.Println("-----------")
	}
}
