package core

import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func StartPacketSniffing(ipv4Only bool, ipv6Only bool){

devices, err := pcap.FindAllDevs()
if err != nil {
	log.Fatalf("Error finding network devices: %v", err)
}

if len(devices) == 0 {
	log.Fatalf("No devices found, make sure you're running as root")
}

device := devices[0].Name

log.Printf("Using device %s\n", device)

handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Error opening device %s: %v", device, err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Monitoring network traffic...")

	for packet := range packetSource.Packets() {
		// Extract network layer info
		netLayer := packet.NetworkLayer()
		if netLayer == nil {
			continue
		}

		// Get source and destination IPs
		srcIP, dstIP := netLayer.NetworkFlow().Endpoints()

		// Apply IPv4 and IPv6 filters
		if ipv4Only && net.ParseIP(srcIP.String()).To4() == nil {
			continue
		}
		if ipv6Only && net.ParseIP(srcIP.String()).To16() != nil && net.ParseIP(srcIP.String()).To4() != nil {
			continue
		}

		fmt.Printf("[Packet] %s -> %s\n", srcIP, dstIP)
	}

}