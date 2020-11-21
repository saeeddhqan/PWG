
package main

import (
	"log"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"time"
	"flag"
	"strings"
)

var iface = flag.String("i", "wlan0", "Network device (default=any)")
var bpf = flag.String("bpf", "tcp and port 80", "Set packet filtering (default='tcp and port 80')")
var snaplen = flag.Int("s", 65535, "Snap length (default=65535)")
var timeout = flag.Int("t", -1, "Deadline for capturing loop (default=-1)")
var promisc = flag.Bool("p", false, "Set promiscuous mode (default=false)")

func main() {
	flag.Parse()
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}
	deviceFound := true
	for _, device := range devices {
		if device.Name == *iface {
			deviceFound = true
		}
	}

	if !deviceFound {
		log.Println("The device name", *iface, "not found!")
		return
	}

	handler, err := pcap.OpenLive(*iface, int32(*snaplen), *promisc, time.Second * time.Duration(*timeout))
	if err != nil {
		log.Fatalln(err)
	}

	if err := handler.SetBPFFilter(*bpf); err != nil {
		log.Fatalln(err)
	}

	for packet := range gopacket.NewPacketSource(handler, handler.LinkType()).Packets() {
		// fmt.Println(packet.Dump())
		data := string(packet.Data())
		if strings.Contains(data, "USER") || strings.Contains(data, "PASS") {
			fmt.Println(data)
			fmt.Println("------------")
		}
	}

}