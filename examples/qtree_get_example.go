package main

import (
	"fmt"
	"go-ontap-sdk/ontap"
	"time"
)

func main() {
	client := ontap.NewClient(
		"https://mycluster.example.com",
		&ontap.ClientOptions{
			Version:           "1.160",
			BasicAuthUser:     "admin",
			BasicAuthPassword: "secret",
			SSLVerify:         false,
			Debug:             false, // Enable debugging
			Timeout:           60 * time.Second,
		},
	)
	qtreeOptions := &ontap.QtreeGetIterOptions{
		MaxRecords: 1024,
	}
	qtreeResponses, err := client.QtreeGetIterAPI(qtreeOptions)
	if err != nil {
		fmt.Printf("Error fetching Qtree information: %v\n", err)
		return
	}

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
			fmt.Printf("Total Qtrees: %d\n", response.Results.NumRecords)
		}
	} else {
		fmt.Println("No Qtrees found.")
	}
}
