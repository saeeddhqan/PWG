
package main

import (
	"fmt"
	"log"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"flag"
	"time"
)

var iface = flag.String("i", "wlan0", "Interface name to read packets from")
var snaplen = flag.Int("s", 65535, "Snap length (The max number of bytes to read from each packets)")
var timeout = flag.Int("t", -1, "Packet timeout (Per second)")
var promisc = flag.Bool("p", false, "Set promiscuous")

func main() {
	flag.Parse()
	handler, err := pcap.OpenLive(*iface, int32(*snaplen), *promisc, time.Second * time.Duration(*timeout))
	if err != nil {
		log.Fatalln(err)
	}
	defer handler.Close()

	for packet := range gopacket.NewPacketSource(handler, handler.LinkType()).Packets() {
		// fmt.Println(packet) // packet.String()
		// fmt.Println(packet.Dump())
		fmt.Println(string(packet.Data()))
		fmt.Println("-----------")
	}
}
