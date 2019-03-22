package main

import (
	"fmt"
	"net/http"
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

// Pings Master Node with updated Sub Node information
func updateSubNodeInfo(ip string, port string) {
	url := "http://" + ip + ":" + port + "/api/update-sub-node?BatteryLevelPercentage=" + getBatteryPercentage() + "&IP=" + getIp()

	resp, err := http.Get(url)
	handleError(err)

	fmt.Printf("Response: %s\n", resp)

	// fmt.Printf("Response: %s\n", resp)
}

func test(ip string, port string) {
	url := "http://" + ip + ":" + port + "/api/update-sub-node?BatteryLevelPercentage=67.89&IP=123.456.789.000"
	resp, err := http.Get(url)
	handleError(err)

	fmt.Printf("Response: %s\n", resp)
}

func main() {
	// Constants
	const googleComputeIp = "35.243.155.9"
	const localIp = "127.0.0.1"
	var port = "3141"

	// resp, err := http.Get("http://35.243.155.9:3141/api/update-sub-node?BatteryLevelPercentage=67.89&IP=123.456.789.000")
	// handleError(err)
	// fmt.Printf("Response: %s\n", resp)

	updateSubNodeInfo(googleComputeIp, port)
	// test(googleComputeIp, port)
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
