/*
Copyright © 2025 @mdxabu
*/
package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the ipscout configuration",
	Long:  `Initialize the ipscout configuration file with Wi-Fi IP detection.`,
	Run: func(cmd *cobra.Command, args []string) {
		configFilePath := "ipscoutconfig.yaml"

		if _, err := os.Stat(configFilePath); err == nil {
			fmt.Printf("Configuration file '%s' already exists.\n", configFilePath)
			return
		}

		fmt.Println("Creating YAML configuration file...")

		interfaces, err := net.Interfaces()
		if err != nil {
			fmt.Println("Error getting network interfaces:", err)
			return
		}

		var wifiInterface *net.Interface
		for _, iface := range interfaces {
			if iface.Flags&net.FlagUp != 0 && (iface.Name == "Wi-Fi" || iface.Name == "wlan0") {
				wifiInterface = &iface
				break
			}
		}

		if wifiInterface == nil {
			fmt.Println("No Wi-Fi interface detected.")
			return
		}

		var ipv4, ipv6 string
		addrs, err := wifiInterface.Addrs()
		if err != nil {
			fmt.Printf("Error getting addresses for interface %s: %v\n", wifiInterface.Name, err)
			return
		}

		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				continue
			}
			if ip.To4() != nil && ipv4 == "" {
				ipv4 = ip.String()
			} else if ip.To16() != nil && ipv6 == "" && ip.To4() == nil {
				ipv6 = ip.String()
			}
		}

		fmt.Printf("Detected Wi-Fi interface: %s\n", wifiInterface.Name)
		fmt.Printf("IPv4: %s\n", ipv4)
		fmt.Printf("IPv6: %s\n", ipv6)

		config := map[string]interface{}{
			"wifi": map[string]interface{}{
				"name": wifiInterface.Name,
				"ipv4": ipv4,
				"ipv6": ipv6,
			},
		}

		configData, err := yaml.Marshal(config)
		if err != nil {
			fmt.Println("Error marshaling YAML:", err)
			return
		}

		err = os.WriteFile(configFilePath, configData, 0644)
		if err != nil {
			fmt.Println("Error writing to config file:", err)
			return
		}

		fmt.Printf("✓ Configuration file '%s' created successfully with Wi-Fi details.\n", configFilePath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
