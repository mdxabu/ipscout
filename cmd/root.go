/*
Copyright © 2024 @mdxabu
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)





var rootCmd = &cobra.Command{
	Use:   "ipscout",
	Short: "ipscout is a tool to get the geo location of the ip address",
	Long: `ipscout is a tool to get the geo location of the ip address.`,

	Run: func(cmd *cobra.Command, args []string) {

		asciibanner := `
██╗██████╗ ███████╗ ██████╗ ██████╗ ██╗   ██╗████████╗
██║██╔══██╗██╔════╝██╔════╝██╔═══██╗██║   ██║╚══██╔══╝
██║██████╔╝███████╗██║     ██║   ██║██║   ██║   ██║   
██║██╔═══╝ ╚════██║██║     ██║   ██║██║   ██║   ██║   
██║██║     ███████║╚██████╗╚██████╔╝╚██████╔╝   ██║   
╚═╝╚═╝     ╚══════╝ ╚═════╝ ╚═════╝  ╚═════╝    ╚═╝   `

		fmt.Println(asciibanner)
        fmt.Println("Welcome to ipscout! Use 'ipscout --help' to see available commands.")
    },
	
  }
  
  func Execute() error {
	return rootCmd.Execute()
}