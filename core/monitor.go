package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func getInterfaceIPs(interfaceName string) (map[string]bool, error) {
	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return nil, err
	}
	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}
	ipMap := make(map[string]bool)
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip != nil {
			ipMap[ip.String()] = true
		}
	}
	return ipMap, nil
}

func resolveHost(ip string) string {
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		return "-"
	}
	return strings.TrimSuffix(names[0], ".")
}

type GeoInfo struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Query   string `json:"query"`
}

var geoCache = make(map[string]GeoInfo)

func getGeoInfo(ip string) GeoInfo {
	if info, ok := geoCache[ip]; ok {
		return info
	}
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return GeoInfo{}
	}
	defer resp.Body.Close()
	var info GeoInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return GeoInfo{}
	}
	geoCache[ip] = info
	return info
}

func StartPacketSniffing(interfaceName string, useIPv4 bool, useIPv6 bool, filterSrcIP string) {
	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Error opening device %s: %v", interfaceName, err)
	}
	defer handle.Close()

	fmt.Printf("Sniffing on interface: %s\n", interfaceName)
	if !useIPv4 && !useIPv6 {
		fmt.Println("Warning: Neither IPv4 nor IPv6 flag set, capturing all packets.")
	}

	_, err = getInterfaceIPs(interfaceName)
	if err != nil {
		fmt.Println("Warning: Could not get local IPs for interface:", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetChan := packetSource.Packets()

	lastPacket := time.Now()
	timeout := 10 * time.Second

	fmt.Printf("%-40s %-16s %-24s %-40s %-16s %-24s %-8s\n", "Sender Name", "Sender IP", "Sender Location", "Receiver Name", "Receiver IP", "Receiver Location", "Proto")
	fmt.Printf("%s\n", strings.Repeat("-", 40+16+24+40+16+24+8+6))

	for {
		select {
		case packet, ok := <-packetChan:
			if !ok {
				fmt.Println("Packet source closed.")
				return
			}
			lastPacket = time.Now()
			networkLayer := packet.NetworkLayer()
			if networkLayer == nil {
				continue
			}

			var srcIP, dstIP string
			var proto string
			var isIPv4, isIPv6, isICMP bool

			if useIPv4 && networkLayer.LayerType() == layers.LayerTypeIPv4 {
				ipv4 := networkLayer.(*layers.IPv4)
				srcIP, dstIP = ipv4.SrcIP.String(), ipv4.DstIP.String()
				isIPv4 = true
			} else if useIPv6 && networkLayer.LayerType() == layers.LayerTypeIPv6 {
				ipv6 := networkLayer.(*layers.IPv6)
				srcIP, dstIP = ipv6.SrcIP.String(), ipv6.DstIP.String()
				isIPv6 = true
			} else if !useIPv4 && !useIPv6 {
				switch l := networkLayer.(type) {
				case *layers.IPv4:
					srcIP, dstIP = l.SrcIP.String(), l.DstIP.String()
					isIPv4 = true
				case *layers.IPv6:
					srcIP, dstIP = l.SrcIP.String(), l.DstIP.String()
					isIPv6 = true
				default:
					continue
				}
			} else {
				continue
			}

			if isIPv4 {
				if icmpLayer := packet.Layer(layers.LayerTypeICMPv4); icmpLayer != nil {
					isICMP = true
					proto = "ICMP"
				}
			} else if isIPv6 {
				if icmpLayer := packet.Layer(layers.LayerTypeICMPv6); icmpLayer != nil {
					isICMP = true
					proto = "ICMP"
				}
			}

			if !isICMP {
				if transportLayer := packet.TransportLayer(); transportLayer != nil {
					switch transportLayer.LayerType() {
					case layers.LayerTypeTCP:
						proto = "TCP"
					case layers.LayerTypeUDP:
						proto = "UDP"
					default:
						proto = transportLayer.LayerType().String()
					}
				} else {
					proto = networkLayer.LayerType().String()
				}
			}

			srcHost := resolveHost(srcIP)
			dstHost := resolveHost(dstIP)

			srcGeo := getGeoInfo(srcIP)
			dstGeo := getGeoInfo(dstIP)
			srcLoc := fmt.Sprintf("%s, %s", srcGeo.Country, srcGeo.City)
			dstLoc := fmt.Sprintf("%s, %s", dstGeo.Country, dstGeo.City)

			fmt.Printf("%-40s %-16s %-24s %-40s %-16s %-24s %-8s\n",
				srcHost, srcIP, srcLoc, dstHost, dstIP, dstLoc, proto,
			)
		default:
			if time.Since(lastPacket) > timeout {
				fmt.Println("No packets captured in the last 10 seconds. Are you using the correct interface? Try running as administrator.")
				lastPacket = time.Now()
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}
