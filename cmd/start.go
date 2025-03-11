/*
Copyright © 2025 @mdxabu
*/
package cmd

import (
	"fmt"
	"github.com/mdxabu/ipscout/core"
	"github.com/spf13/cobra"	
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Command to start the ipscout process",
	Long: `Command to start the ipscout process. This command will start the ipscout process and begin monitoring the network.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Process started, This may take a few seconds to start monitoring the network.")
		// IPv4, IPv6 , err := core.FetchYaml()
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// useIPv4, _ := cmd.Flags().GetBool("ipv4")
		// useIPv6, _ := cmd.Flags().GetBool("ipv6")
		useMonitor, _ := cmd.Flags().GetBool("monitor")

		if useMonitor {
			useIPv6, _ := cmd.Flags().GetBool("ipv6")
			useIPv4, _ := cmd.Flags().GetBool("ipv4")

			core.StartPacketSniffing(useIPv4, useIPv6)
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
