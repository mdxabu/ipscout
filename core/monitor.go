package core

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/layers"
)

// StartPacketSniffing captures packets on the specified interface
func StartPacketSniffing(interfaceName string, useIPv4 bool, useIPv6 bool) {
	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Error opening device %s: %v", interfaceName, err)
	}
	defer handle.Close()

	fmt.Printf("Sniffing on interface: %s\n", interfaceName)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		networkLayer := packet.NetworkLayer()
		if networkLayer == nil {
			continue
		}

		if useIPv4 && networkLayer.LayerType() == layers.LayerTypeIPv4 {
			ipv4 := networkLayer.(*layers.IPv4)
			src, dst := ipv4.SrcIP, ipv4.DstIP
			fmt.Printf("IPv4 Packet: %s -> %s\n", src, dst)
		} else if useIPv6 && networkLayer.LayerType() == layers.LayerTypeIPv6 {
			ipv6 := networkLayer.(*layers.IPv6)
			src, dst := ipv6.SrcIP, ipv6.DstIP
			fmt.Printf("IPv6 Packet: %s -> %s\n", src, dst)
		}
		}
	}

