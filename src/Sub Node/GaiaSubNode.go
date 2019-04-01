package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// Gets networking port information
func getOpenPorts() string {
	cmd := "./Bash Functions/getOpenPorts.sh"

	// Get's output of 'nmap' command
	openPortsByte, _ := exec.Command(cmd).Output()
	openPortsString := string(openPortsByte)
	openPortsString = strings.Trim(openPortsString, "\n")

	return openPortsString
}

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

	// Gets battery level
	batteryLevelPercentageBytes, err := exec.Command(cmd).Output()
	handleError(err)

	// Trims '\n' from end of string
	batteryLevelPercentage := string(batteryLevelPercentageBytes)
	batteryLevelPercentage = strings.TrimSuffix(batteryLevelPercentage, "\n")

	return batteryLevelPercentage
}

// Takes screenshot of current screen
func takeScreenshot() {
	cmd := "./Bash Functions/getScreenshot.sh"
	_, err := exec.Command(cmd).Output()
	handleError(err)
}

// Takes screenshot of desktop, then returns a base64 string representation of the screenshot
func getBase64Screenshot() string {
	// Takes screenshot
	takeScreenshot()

	// Loads screenshot image into memory as base64 string
	imageName := "Screenshots/screenshot.png"
	buf, err := ioutil.ReadFile(imageName)
	if err != nil {
		log.Fatal(err)
	}
	encoded := base64.StdEncoding.EncodeToString(buf)

	return encoded
}

// Pings Master Node with system information
func pingMasterNode(url string, subNodeIP string, updateTimeDelaySecond float64) {
	// Gets system information and adds it to POST request payload
	base64Image := getBase64Screenshot()
	batteryLevelPercentage := getBatteryPercentage()
	portInfo := getOpenPorts()
	fmt.Println(portInfo)
	payload := strings.NewReader(fmt.Sprintf("Base64Image=%s&BatteryLevelPercentage=%s&IP=%s&PortInfo=%s", base64Image, batteryLevelPercentage, subNodeIP, portInfo))

	// Defines POST request with appropriate header(s)
	req, err := http.NewRequest("POST", url, payload)
	handleError(err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")                             // ??? Why is this required ???
	req.Header.Add("Postman-Token", "93b74e57-31fd-436d-976c-e10f223dc185") // ??? Why is this required ???

	// Sends POST request
	res, err := http.DefaultClient.Do(req)
	handleError(err)
	respBody, err := ioutil.ReadAll(res.Body)
	handleError(err)
	fmt.Println(string(respBody))
	res.Body.Close()
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Master Node IP and port
	const googleMasterNodeIP = "http://35.243.155.9"
	const localMasterNodeIP = "http://127.0.0.1"

	// Configures network
	var masterNodeIP = googleMasterNodeIP
	var masterNodePort = "3141" // Port on which Master Node listens
	pingEndpoint := "/api/ping" // Ping API endpoint

	// Time delay between pings to Master Node
	var updateTimeDelaySecond float64 = 5
	timedelayMilliseconds := time.Duration(updateTimeDelaySecond * 1000)

	// Listens for instructions from Master Node
	// TODO

	// Constructs url to Master Node
	url := masterNodeIP + ":" + masterNodePort + pingEndpoint
	// Gets Sub Node IP
	subNodeIP := getSubNodeIP()

	// Pings Master Node at regular intervals with new system information
	for {
		pingMasterNode(url, subNodeIP, updateTimeDelaySecond)
		time.Sleep(timedelayMilliseconds * time.Millisecond)
	}
}
