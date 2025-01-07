/*
Copyright © 2025 @mdxabu
*/
package cmd

import (
	"fmt"
	"github.com/ip2location/ip2location-go/v9"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

// Helper function to validate IP address range
func validateIPRange(ip string) bool {
	// Regular expression to validate the format of the IP address
	// Checks if the first part of the IP is between 0 and 99, and each part is in the range of 0-255
	re := regexp.MustCompile(`^([0-9]{1,2}|[0-9]{1,2}[0-9]{1,2})\.(?:[0-9]{1,3})\.(?:[0-9]{1,3})\.(?:[0-9]{1,3})$`)
	if !re.MatchString(ip) {
		return false
	}
	// Split the IP into its components
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}
	// Validate that the first part is in the range of 0-99
	firstOctet := parts[0]
	if firstOctet == "" || len(firstOctet) > 2 || firstOctet > "99" {
		return false
	}

	return true
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Finding the location of the ip address",
	Long:  `Finding the geolocation of the ip address.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure an IP address is provided
		if len(args) < 1 {
			fmt.Println("Please provide an IP address")
			return
		}
		ip := args[0]

		// Validate the IP address range
		if !validateIPRange(ip) {
			fmt.Println("Invalid IP address. The first part of the IP address must be between 0-99.")
			return
		}

		// Open the database
		db, err := ip2location.OpenDB("./DBIPV4/IPV4LOCATION.BIN")
		if err != nil {
			fmt.Println("Error opening database:", err)
			return
		}
		defer db.Close()

		// Get the geolocation data
		results, err := db.Get_all(ip)
		if err != nil {
			fmt.Println("Error retrieving data:", err)
			return
		}

		// Display the results
		fmt.Printf("country_short: %s\n", results.Country_short)
		fmt.Printf("country_long: %s\n", results.Country_long)
		fmt.Printf("region: %s\n", results.Region)
		fmt.Printf("city: %s\n", results.City)
		fmt.Printf("isp: %s\n", results.Isp)
		fmt.Printf("latitude: %f\n", results.Latitude)
		fmt.Printf("longitude: %f\n", results.Longitude)
		fmt.Printf("domain: %s\n", results.Domain)
		fmt.Printf("zipcode: %s\n", results.Zipcode)
		fmt.Printf("timezone: %s\n", results.Timezone)
		fmt.Printf("netspeed: %s\n", results.Netspeed)
		fmt.Printf("iddcode: %s\n", results.Iddcode)
		fmt.Printf("areacode: %s\n", results.Areacode)
		fmt.Printf("weatherstationcode: %s\n", results.Weatherstationcode)
		fmt.Printf("weatherstationname: %s\n", results.Weatherstationname)
		fmt.Printf("mcc: %s\n", results.Mcc)
		fmt.Printf("mnc: %s\n", results.Mnc)
		fmt.Printf("mobilebrand: %s\n", results.Mobilebrand)
		fmt.Printf("elevation: %f\n", results.Elevation)
		fmt.Printf("usagetype: %s\n", results.Usagetype)
		fmt.Printf("addresstype: %s\n", results.Addresstype)
		fmt.Printf("category: %s\n", results.Category)
		fmt.Printf("district: %s\n", results.District)
		fmt.Printf("asn: %s\n", results.Asn)
		fmt.Printf("as: %s\n", results.As)
		fmt.Printf("api version: %s\n", ip2location.Api_version())
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}
