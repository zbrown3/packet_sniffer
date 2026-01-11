package main
import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)


func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}

	device := devices[0].Name

	handle, err := pcap.OpenLive(device, 65536, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}

	defer handle.Close()
	
	fmt.Printf("Listening on interface %s...\n", device)
} 
