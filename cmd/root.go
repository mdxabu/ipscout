/*
Copyright © 2024 @mdxabu
*/
package cmd

import (

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "ipscout",
	Short: "ipscout is a tool to get the geo location of the ip address",
	Long: `ipscout is a tool to get the geo location of the ip address.`,
  }
  
  func Execute() error {
	return rootCmd.Execute()
}