package main

import (
	"fmt"
	"go-ontap-sdk/ontap"
	"time"
)

func main() {
	// Initialize the ONTAP client
	client := ontap.NewClient(
		"http://192.168.200.60",
		&ontap.ClientOptions{
			Version:           "1.32", // Ensure this version is compatible with your ONTAP system
			BasicAuthUser:     "umonitor",
			BasicAuthPassword: "sxZeJs4n",
			SSLVerify:         false,
			Debug:             false, // Enable debugging
			Timeout:           60 * time.Second,
		},
	)

	// Fetch Qtree information iteratively
	qtreeOptions := &ontap.QtreeListIterOptions{
		MaxRecords: 1024,
	}

	var allQtrees []*ontap.QtreeListIterResponse
	qtreeResponses, err := client.QtreeListIterAPI(qtreeOptions)
	if err != nil {
		fmt.Printf("Error fetching Qtree information: %v\n", err)
		return
	}

	// Process and display the Qtree information
	if len(qtreeResponses) > 0 {
		fmt.Println("Qtree Information:")
		for _, response := range qtreeResponses {
			for _, qtree := range response.Results.AttributesList.Qtrees {
				fmt.Printf("Qtree: %s\n", qtree.Qtree)
				fmt.Printf("Volume: %s\n", qtree.Volume)
				fmt.Printf("Vserver: %s\n", qtree.Vserver)
				fmt.Printf("Status: %s\n", qtree.Status)
				fmt.Printf("OpLocks: %s\n", qtree.OpLocks)
				fmt.Printf("Security Style: %s\n", qtree.SecurityStyle)
				fmt.Println()
			}
			allQtrees = append(allQtrees, response)
			fmt.Printf("Total Qtrees: %d\n", response.Results.NumRecords)

		}
		//fmt.Printf("Total Qtrees: %d\n", len(allQtrees))
	} else {
		fmt.Println("No Qtrees found.")
	}
}
