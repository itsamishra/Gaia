package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Returns external IP associated with this machine
func getSubNodeIP() string {
	resp, err := http.Get("http://ifconfig.me")
	handleError(err)
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	defer resp.Body.Close()

	return string(body)
}

// Returns battery level (%) of this machine
func getBatteryPercentage() string {
	// Bash file containing command that gets battery level
	cmd := "./Bash Functions/getBatteryLevel.sh"
	// cmd := "/home/am/go/src/Gaia/src/Sub Node/Bash Functions/getBatteryLevel.sh"
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
	// // TODO: Take screenshot (use command: import -window root root.png)

	// Reads image file
	// imageName := "Screenshots/screenshot.png"
	imageName := "Screenshots/screen.png"
	// imageName := "/home/am/go/src/Gaia/src/Sub Node/Screenshots/screenshot.png"

	// imageFileBytes, err := ioutil.ReadFile(imageName)
	// handleError(err)

	buf := bytes.NewBuffer(nil)
	f, err := os.Open(imageName)
	handleError(err)

	io.Copy(buf, f)
	fmt.Println("okay")
	handleError(err)
	f.Close()

	// Converts file to base64 string
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	return encoded
	// imageName := "Screenshots/s-shot.png"
	// buf, err := ioutil.ReadFile(imageName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// encoded := base64.StdEncoding.EncodeToString(buf)
	// return encoded
}

// Pings Master Node with updated Sub Node information every 'n' seconds
func updateSubNodeInfo(masterNodeIP string, port string, updateTimeDelaySecond float64) {
	// timedelayMilliseconds := time.Duration(updateTimeDelaySecond * 1000)

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
		// time.Sleep(timedelayMilliseconds * time.Millisecond)
	}
}

func pingMasterNode(masterNodeIP string, port string, updateTimeDelaySecond float64, endpoint string) {
	// Converts time delay from seconds to milliseconds
	timedelayMilliseconds := time.Duration(updateTimeDelaySecond * 1000)

	// Constructs url to Master Node
	url := masterNodeIP + ":" + port + endpoint
	subNodeIP := getSubNodeIP()
	for {
		// Gets system information and adds it to POST request payload
		// TODO: Add screenshot capability
		// base64Image := getBase64Screenshot()
		base64Image := "TestImagePleaseIgnore"
		batteryLevelPercentage := getBatteryPercentage()
		payload := strings.NewReader(fmt.Sprintf("Base64Image=%s&BatteryLevelPercentage=%s&IP=%s", base64Image, batteryLevelPercentage, subNodeIP))

		// Defines POST request with appropriate header(s)
		req, err := http.NewRequest("POST", url, payload)
		handleError(err)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		// req.Header.Add("cache-control", "no-cache")
		// req.Header.Add("Postman-Token", "93b74e57-31fd-436d-976c-e10f223dc185")

		// Sends POST request
		res, err := http.DefaultClient.Do(req)
		handleError(err)
		body, err := ioutil.ReadAll(res.Body)
		handleError(err)
		fmt.Println(string(body))
		res.Body.Close()

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
	// var masterNodeIP = "http://35.243.155.9"
	var localMasterNodeIP = "http://127.0.0.1"

	// Port on which Master Node listens
	var masterNodePort = "3141"
	// Time delay between pings to Master Node
	var updateTimeDelaySecond float64 = 1
	// Ping API endpoint
	pingEndpoint := "/api/ping"

	// Constantly pings Master Node with updated Sub Node information
	go pingMasterNode(localMasterNodeIP, masterNodePort, updateTimeDelaySecond, pingEndpoint)

	// Infinite loop
	for {
	}

	// Listens for instructions from Master Node
	// TODO
}
