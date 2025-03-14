package core

import (
	"fmt"
	"log"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func StartPacketSniffing(useIPv4 bool, useIPv6 bool) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("Error finding network devices: %v", err)
	}

	if len(devices) == 0 {
		log.Println("No network devices found.")
		return
	}

	device := devices[0].Name
	fmt.Printf("Using device: %s\n", device)

	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Error opening device %s: %v", device, err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
