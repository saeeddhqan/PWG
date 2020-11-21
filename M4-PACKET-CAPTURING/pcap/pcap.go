
package main

import (
	"log"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"time"
	"flag"
	"os"
)

var iface = flag.String("i", "wlan0", "Network device (default=any)")
var bpf = flag.String("bpf", "", "Set packet filtering (default='tcp and port 80')")
var fileName = flag.String("f", "Capture", "PCAP file name (default=Capture)")
var numberOfPackets = flag.Int("c", 50, "Maximum number of packets (default=50")
var snaplen = flag.Int("s", 65535, "Snap length (default=65535)")
var timeout = flag.Int("t", -1, "Deadline for capturing loop (default=-1)")
var promisc = flag.Bool("p", false, "Set promiscuous mode (default=false)")

func main() {
	flag.Parse()
	packetCount := 0
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

	file, err := os.Create(*fileName + ".pcap")
	if err != nil {
		log.Fatalln(err)
	}
	writer := pcapgo.NewWriter(file)
	writer.WriteFileHeader(uint32(*snaplen), handler.LinkType())
	defer file.Close()

	if err := handler.SetBPFFilter(*bpf); err != nil {
		log.Fatalln(err)
	}

	for packet := range gopacket.NewPacketSource(handler, handler.LinkType()).Packets() {
		fmt.Println(packet)
		writer.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++
		if packetCount == *numberOfPackets {
			break
		}
	}

}