/*
Copyright © 2024 @mdxabu
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the ipscout configuration",
	Long: `Initialize the ipscout configuration. This command creates a configuration file in the user's desired directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating the configuration file...")

		// Define the configuration structure
		config := map[string]interface{}{
			"ipv4": map[string]interface{}{
				"enabled": true,
				"prefix":  "//IPv4 address",
			},
			"ipv6": map[string]interface{}{
				"enabled": false,
				"prefix":  "//IPv6 address",
			},
			"logging": map[string]interface{}{
				"level":  "info",
				"format": "json",
			},
		}

		// Marshal the config to JSON
		configData, err := json.MarshalIndent(config, "", "\t")
		if err != nil {
			fmt.Println("Error marshalling config:", err)
			os.Exit(1)
		}

		// Create or overwrite the configuration file
		configFile, err := os.Create("ipscoutconfig.json")
		if err != nil {
			fmt.Println("Error creating config file:", err)
			os.Exit(1)
		}
		defer configFile.Close()

		// Write the JSON data to the file
		_, err = configFile.Write(configData)
		if err != nil {
			fmt.Println("Error writing to config file:", err)
			os.Exit(1)
		}

		fmt.Println("✓ Configuration file created successfully.")
	},
}


func init() {
	rootCmd.AddCommand(initCmd)
}
