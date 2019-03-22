package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Returns external IP associated with this machine
func getSubNodeIP() string {
	resp, err := http.Get("http://ifconfig.me")
	if err != nil {
		// handle err
	}
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	defer resp.Body.Close()

	return string(body)
}

// Returns battery level (%) of this machine
func getBatteryPercentage() string {
	// TODO: Add functionality
	// User upower -i /org/freedesktop/UPower/devices/battery_BAT0 (works on Ubuntu)
	return "67.89"
}

// TODO get screenshot of machine
func getScreenshot() {

}

// Pings Master Node with updated Sub Node information every 'n' seconds
func updateSubNodeInfo(masterNodeIP string, port string, updateTimeDelaySecond float64) {
	timedelayMilliseconds := time.Duration(updateTimeDelaySecond * 1000)

	subNodeIP := getSubNodeIP()
	for {
		// Sends GET request to Master Node with Sub Node information
		url := "http://" + masterNodeIP + ":" + port + "/api/update-sub-node?BatteryLevelPercentage=" + getBatteryPercentage() + "&IP=" + subNodeIP
		// TODO: Send data as POST instead of GET
		resp, err := http.Get(url)
		handleError(err)

		// Prints out message from Master Node
		body, err := ioutil.ReadAll(resp.Body)
		handleError(err)
		fmt.Println(string(body))
		resp.Body.Close()

		// Waits before sending information again
		time.Sleep(timedelayMilliseconds * time.Millisecond)
	}
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Master Node IP and port
	var masterNodeIP = "35.243.155.9"
	var masterNodePort = "3141"
	var updateTimeDelaySecond float64 = 1

	// Constantly updates Master Node with updated Sub Node information
	go updateSubNodeInfo(masterNodeIP, masterNodePort, updateTimeDelaySecond)

	// Infinite loop
	for {
	}

	// Listens for instructions from Master Node
	// TODO
}
