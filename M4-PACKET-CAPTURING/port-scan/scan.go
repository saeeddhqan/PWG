
package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"net"
	"sync"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	iface = "any"
	ip = "45.33.32.156"
	ports = strings.Split("21,22,25,80,135,1723,3306,9929", ",")
	bpf = "src host %s and (tcp[13] == 0x11 or tcp[13] == 0x10 or tcp[13] == 0x18)"
	snaplen = int32(1024)
	timeout = pcap.BlockForever
	promisc = false
	response = make(map[string]int)
)


func main() {
	bpf = fmt.Sprintf(bpf, ip)

	go func() {
		handler, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
		if err != nil {
			log.Fatalln(err)
		}
		defer handler.Close()

		if err := handler.SetBPFFilter(bpf); err != nil {
			log.Fatalln(err)
		}

		for packet := range gopacket.NewPacketSource(handler, handler.LinkType()).Packets() {
			srcPort := packet.TransportLayer().TransportFlow().Src().String()
			response[srcPort] += 1
			// log.Println(string(packet.Data()))
		}
	}()
	time.Sleep(1 * time.Second)
	var WG sync.WaitGroup
	for _,port := range ports {
		WG.Add(1)
		go func(i string){
			req, err := net.Dial("tcp", ip + ":" + i)
			if err != nil {
				return
			}
			req.Close()
		}(port)
	}

	time.Sleep(time.Second * 7)
	for port, confidence := range response {
		fmt.Printf("Port %s is open, confidence= %d\n", port, confidence)
	}

}