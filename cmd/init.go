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
		configFilePath := "ipscoutconfig.json"

		// Check if the configuration file already exists
		if _, err := os.Stat(configFilePath); err == nil {
			fmt.Printf("Configuration file '%s' already exists.\n", configFilePath)
			return
		} else if !os.IsNotExist(err) {
			// If the error is not due to the file not existing, log it
			fmt.Println("Error checking configuration file:", err)
			return
		}

		// Proceed to create the configuration file
		fmt.Println("Creating the configuration file...")

		// Define the configuration structure
		config := map[string]interface{}{
			"ipv4": map[string]interface{}{
				"enabled": true,
				"prefix":  "// IPv4 Address",
			},
			"ipv6": map[string]interface{}{
				"enabled": false,
				"prefix":  "// IPv6 Address",
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
			return
		}

		// Create the configuration file
		configFile, err := os.Create(configFilePath)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			return
		}
		defer configFile.Close()

		// Write the JSON data to the file
		if _, err := configFile.Write(configData); err != nil {
			fmt.Println("Error writing to config file:", err)
			return
		}

		fmt.Printf("✓ Configuration file '%s' created successfully.\n", configFilePath)
	},
}



func init() {
	rootCmd.AddCommand(initCmd)
}
