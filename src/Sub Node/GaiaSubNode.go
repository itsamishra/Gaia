package main

import (
	"fmt"
	"net/http"
	"time"
)

// Returns external IP associated with this machine
func getIp() string {
	// TODO: Add functionality
	return "123.456.789.000"
}

// Returns battery level (%) of this machine
func getBatteryPercentage() string {
	// TODO: Add functionality
	return "67.89"
}

// Pings Master Node with updated Sub Node information every 'n' seconds
func updateSubNodeInfo(ip string, port string) {
	for {
		url := "http://" + ip + ":" + port + "/api/update-sub-node?BatteryLevelPercentage=" + getBatteryPercentage() + "&IP=" + getIp()

		resp, err := http.Get(url)
		handleError(err)

		fmt.Printf("Response: %s\n", resp)

		time.Sleep(2 * time.Second)
	}
}

func main() {
	// Constants
	const googleComputeIp = "35.243.155.9"
	const localIp = "127.0.0.1"
	var port = "3141"

	// Constantly updates Master Node with updated Sub Node information
	go updateSubNodeInfo(googleComputeIp, port)

	time.Sleep(10 * time.Second)
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
