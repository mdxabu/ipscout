/*
Copyright © 2025 @mdxabu
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Command to start the ipscout process",
	Long: `Command to start the ipscout process. This command will start the ipscout process and begin monitoring the network.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Process started, This may take a few seconds to start monitoring the network.")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolP("monitor", "m", false, "monitor the network")
	startCmd.Flags().BoolP("output", "d", false, "output the network data")
	startCmd.Flags().BoolP("ipv4", "4", false, "output the IPv4 data")
	startCmd.Flags().BoolP("ipv6", "6", false, "output the IPv6 data")
}
