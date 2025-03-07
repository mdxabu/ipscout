package cmd

import (
	"fmt"
	"net"
	"os"
	"gopkg.in/yaml.v3"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the ipscout configuration",
	Long:  `Initialize the ipscout configuration file with auto IP detection.`,
	Run: func(cmd *cobra.Command, args []string) {
		configFilePath := "ipscoutconfig.yaml"

		if _, err := os.Stat(configFilePath); err == nil {
			fmt.Printf("Configuration file '%s' already exists.\n", configFilePath)
			return
		}

		fmt.Println("Creating YAML configuration file...")

		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println("Error getting hostname:", err)
			return
		}

		addresses, _ := net.LookupIP(hostname)

		interfaces, err := net.Interfaces()
		if err != nil {
			fmt.Println("Error getting network interfaces:", err)
			return
		}

		ipv4 := ""
		ipv6 := ""

		for _, addr := range addresses {
			if addr.To4() != nil {
				ipv4 = addr.String()
			} else if addr.To16() != nil {
				ipv6 = addr.String()
			}
		}

		config := map[string]interface{}{
			"hostname": hostname,
			"ipv4": map[string]interface{}{
				"enabled": ipv4 != "",
				"address": ipv4,
			},
			"ipv6": map[string]interface{}{
				"enabled": ipv6 != "",
				"address": ipv6,
			},
			"network_interfaces": []map[string]interface{}{},
		}

		for _, iface := range interfaces {
			addrs, _ := iface.Addrs()
			interfaceInfo := map[string]interface{}{
				"name":      iface.Name,
				"addresses": []string{},
			}
			for _, addr := range addrs {
				interfaceInfo["addresses"] = append(interfaceInfo["addresses"].([]string), addr.String())
			}
			config["network_interfaces"] = append(config["network_interfaces"].([]map[string]interface{}), interfaceInfo)
		}

		configData, err := yaml.Marshal(config)
		if err != nil {
			fmt.Println("Error marshaling YAML:", err)
			return
		}

		configFile, err := os.Create(configFilePath)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			return
		}
		defer configFile.Close()

		if _, err := configFile.Write(configData); err != nil {
			fmt.Println("Error writing to config file:", err)
			return
		}

		fmt.Printf("✓ Configuration file '%s' created with IP addresses.\n", configFilePath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
