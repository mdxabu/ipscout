/*
Copyright © 2025 @mdxabu
*/
package cmd

import (
	"fmt"
	"github.com/mdxabu/ipscout/core"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Command to start the ipscout process",
	Long:  `Command to start the ipscout process. This command will start the ipscout process and begin monitoring the network.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Process started, This may take a few seconds to start monitoring the network.")

		useMonitor, _ := cmd.Flags().GetBool("monitor")
		useIPv4, _ := cmd.Flags().GetBool("ipv4")
		useIPv6, _ := cmd.Flags().GetBool("ipv6")

		if useMonitor {
			fmt.Println("Starting packet sniffing...")
			
			// Get active network interfaces
			devices, err := pcap.FindAllDevs()
			if err != nil {
				fmt.Println("Error finding network devices:", err)
				return
			}

			var activeDevice string
			for _, device := range devices {
				if len(device.Addresses) > 0 {
					activeDevice = device.Name
					break
				}
			}

			if activeDevice == "" {
				fmt.Println("No active network interfaces found.")
				return
			}

			fmt.Println("Using device:", activeDevice)
			core.StartPacketSniffing(activeDevice, useIPv4, useIPv6)
		} else {
			fmt.Println("Monitoring not enabled. Use --monitor flag to start packet sniffing.")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolP("monitor", "m", false, "monitor the network")
	startCmd.Flags().BoolP("output", "d", false, "output the network data")
	startCmd.Flags().BoolP("ipv4", "4", false, "output the IPv4 data")
	startCmd.Flags().BoolP("ipv6", "6", false, "output the IPv6 data")
}
