package core

import (
	"encoding/json"
	"fmt"
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

func printTableHeader(ipv6 bool) {
	if ipv6 {
		LogNetworkEventHeaderIPv6()
	} else {
		LogNetworkEventHeaderIPv4()
	}
}

func printTableRow(senderName, senderIP, senderLoc, receiverName, receiverIP, receiverLoc, proto string, ipv6 bool) {
	if ipv6 {
		LogNetworkEventIPv6(senderName, senderIP, senderLoc, receiverName, receiverIP, receiverLoc, proto)
	} else {
		LogNetworkEventIPv4(senderName, senderIP, senderLoc, receiverName, receiverIP, receiverLoc, proto)
	}
}

func StartPacketSniffing(interfaceName string, useIPv4 bool, useIPv6 bool, filterSrcIP string) {
	Info("Sniffing on interface: %s", interfaceName)
	if !useIPv4 && !useIPv6 {
		Warn("Neither IPv4 nor IPv6 flag set, capturing all packets.")
	}

	packetSource := gopacket.NewPacketSource(
		func() *pcap.Handle {
			handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
			if err != nil {
				Error("Error opening device %s: %v", interfaceName, err)
				return nil
			}
			return handle
		}(),
		layers.LinkTypeEthernet,
	)
	packetChan := packetSource.Packets()

	lastPacket := time.Now()
	timeout := 10 * time.Second

	headerPrintedV4 := false
	headerPrintedV6 := false

	for {
		select {
		case packet, ok := <-packetChan:
			if !ok {
				Warn("Packet source closed.")
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
				if !headerPrintedV4 {
					printTableHeader(false)
					headerPrintedV4 = true
				}
				if icmpLayer := packet.Layer(layers.LayerTypeICMPv4); icmpLayer != nil {
					isICMP = true
					proto = "ICMP"
				}
			} else if isIPv6 {
				if !headerPrintedV6 {
					printTableHeader(true)
					headerPrintedV6 = true
				}
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

			if isIPv6 {
				printTableRow(srcHost, srcIP, srcLoc, dstHost, dstIP, dstLoc, proto, true)
			} else {
				printTableRow(srcHost, srcIP, srcLoc, dstHost, dstIP, dstLoc, proto, false)
			}
		default:
			if time.Since(lastPacket) > timeout {
				Warn("No packets captured in the last 10 seconds. Are you using the correct interface? Try running as administrator.")
				lastPacket = time.Now()
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}
