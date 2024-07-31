package main

import (
	"fmt"
	"go-ontap-sdk/ontap"
	"time"
)

func main() {
	c := ontap.NewClient(
		"https://mycluster.example.com",
		&ontap.ClientOptions{
			Version:           "1.160",
			BasicAuthUser:     "admin",
			BasicAuthPassword: "secret",
			SSLVerify:         false,
			Debug:             false,
			Timeout:           60 * time.Second,
		},
	)

	// Fetch Qtree information
	qtreeResponse, err := c.QtreeGetIter()
	if err != nil {
		fmt.Printf("Error fetching Qtree information: %v\n", err)
		return
	}

	if len(qtreeResponse.Results.AttributesList.Qtrees) > 0 {
		for _, qtree := range qtreeResponse.Results.AttributesList.Qtrees {
			fmt.Printf("Qtree: %s\n", qtree.Qtree)
			fmt.Printf("Volume: %s\n", qtree.Volume)
			fmt.Printf("Vserver: %s\n", qtree.Vserver)
			fmt.Printf("Status: %s\n", qtree.Status)
			fmt.Printf("OpLocks: %s\n", qtree.OpLocks)
			//	fmt.Printf("Security Style: %s\n", qtree.SecurityStyle)
			fmt.Println()
		}
		fmt.Printf("Total Qtrees: %d\n", len(qtreeResponse.Results.AttributesList.Qtrees))
	} else {
		fmt.Println("No Qtrees found.")
	}
}
