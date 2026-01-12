package main
import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)


func main() {

	// Locates all NICs
	devices, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}
	
	// Assigns first NIC to device variable
	device := devices[0].Name
	
	// Opens the interface for packet capture. Sets the max packet size.  Sets promiscuous mode to true.  Sets duration to indefinite.
	handle, err := pcap.OpenLive(device, 65536, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}
	
	// Creates a packet source and iterates over the channel and process each packet.
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}

	defer handle.Close()
	
	fmt.Printf("Listening on interface %s...\n", device)
} 
