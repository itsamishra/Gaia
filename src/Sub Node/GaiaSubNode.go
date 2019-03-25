package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
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

// Takes screenshot of desktop, then returns a base64 string representation of the screenshot
func getBase64Screenshot() string {
	// TODO: Take screenshot (use command: import -window root root.png)

	// Reads image file
	imageName := "Screenshots/screenshot.png"
	imageFileBytes, err := ioutil.ReadFile(imageName)
	if err != nil {
		panic(err)
	}
	fmt.Println("File Size:")
	fmt.Println(strconv.Itoa(len(imageFileBytes)) + " bytes")

	// Converts file to base64 string
	encoded := base64.StdEncoding.EncodeToString(imageFileBytes)
	fmt.Println(len(encoded))
	// fmt.Println(encoded)

	return encoded
}

// Pings Master Node with updated Sub Node information every 'n' seconds
func updateSubNodeInfo(masterNodeIP string, port string, updateTimeDelaySecond float64) {
	timedelayMilliseconds := time.Duration(updateTimeDelaySecond * 1000)

	subNodeIP := getSubNodeIP()
	for {
		// Sends GET request to Master Node with Sub Node information
		url := "http://" + masterNodeIP + ":" + port + "/api/update-sub-node?BatteryLevelPercentage=" + getBatteryPercentage() + "&IP=" + subNodeIP
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
	url := "http://127.0.0.1:3141/api/ping"

	base64Image := getBase64Screenshot()
	// TODO: Replace 56.78 with actual battery %
	batteryLevelPercentage := "56.78"
	ip := getSubNodeIP()
	payload := strings.NewReader(fmt.Sprintf("Base64Image=%s&BatteryLevelPercentage=%s&IP=%s", base64Image, batteryLevelPercentage, ip))

	req, err := http.NewRequest("POST", url, payload)
	handleError(err)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("cache-control", "no-cache")
	// req.Header.Add("Postman-Token", "93b74e57-31fd-436d-976c-e10f223dc185")

	res, err := http.DefaultClient.Do(req)
	handleError(err)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	handleError(err)

	// fmt.Println(res)
	fmt.Println(string(body))
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
	getBase64Screenshot()
	pingMasterNode(localMasterNodeIP, masterNodePort, updateTimeDelaySecond)

	// Infinite loop
	// for {
	// }

	// Listens for instructions from Master Node
	// TODO
}
