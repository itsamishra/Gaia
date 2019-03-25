package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
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
	// Bash file containing command that gets battery level
	cmd := "./Bash Functions/getBatteryLevel.sh"

	// Gets battery level
	batteryLevelPercentageBytes, err := exec.Command(cmd).Output()
	handleError(err)

	// Trims \n from end of string
	batteryLevelPercentage := string(batteryLevelPercentageBytes)
	batteryLevelPercentage = strings.TrimSuffix(batteryLevelPercentage, "\n")

	return batteryLevelPercentage
}

// TODO get screenshot of machine
// Command (Ubuntu): import -window root root.png
func getScreenshot() {

}

func loadImage(fileName string) {
	fmt.Println(fileName)
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

func pingMasterNode(masterNodeIP string, port string, updateTimeDelaySecond float64) {
	url := "http://" + masterNodeIP + ":" + port + "/api/ping"
	fmt.Println(url)

	jsonString := []byte(`{"Name": "John","Photo": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEUAAACnej3aAAAAAXRSTlMAQObYZgAAAApJREFUCNdjYAAAAAIAAeIhvDMAAAAASUVORK5CYII="}`)
	_, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	handleError(err)
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Master Node IP and port
	// var masterNodeIP = "35.243.155.9"
	var localMasterNodeIP = "127.0.0.1"
	var masterNodePort = "3141"
	var updateTimeDelaySecond float64 = 1

	// Constantly updates Master Node with updated Sub Node information
	// go updateSubNodeInfo(localMasterNodeIP, masterNodePort, updateTimeDelaySecond)

	pingMasterNode(localMasterNodeIP, masterNodePort, updateTimeDelaySecond)

	// Infinite loop
	for {
	}

	// Listens for instructions from Master Node
	// TODO
}
