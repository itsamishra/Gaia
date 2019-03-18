package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Message received
type MasterNodeMessage struct {
	WantResponse               bool
	WantSystemInfo             bool
	TimeDelayOnResponseSeconds int
	Timestamp                  int64 // When was this sent?
}

// Message sent
type SlaveNodeMessage struct {
	IsPollMessage bool // Is this message polling the server?
	IsDataMessage bool // Is this message returning data?
	SystemInfo    SystemInfo
	Timestamp     int64 // When was this sent?
}

// Struct containing all important system information
type SystemInfo struct {
	SystemUser         string
	SystemType         string
	SystemHostname     string
	BatteryLifePercent float64
}

// Returns the current battery %
func getBatteryPercentageMacOS() float64 {
	// Runs bash command that gets battery % and converts battery % from string to float64
	commandString := [3]string{"bash", "-c", `pmset -g batt | grep -Eo "\d+%" | cut -d% -f1`}

	command := exec.Command(
		commandString[0],
		commandString[1],
		commandString[2],
	)
	rawOutput, _ := command.CombinedOutput()
	rawOutputString := strings.TrimSuffix(string(rawOutput), "\n")

	batteryPercentage, _ := strconv.ParseFloat(rawOutputString, 64)

	return batteryPercentage
}

// Returns the system id
func getSystemUserMacOS() string {
	commandString := [3]string{"bash", "-c", `whoami`}

	command := exec.Command(
		commandString[0],
		commandString[1],
		commandString[2],
	)
	rawOutput, _ := command.CombinedOutput()
	rawOutputString := strings.TrimSuffix(string(rawOutput), "\n")

	return rawOutputString
}

func getSystemHostnameMacOS() string {
	commandString := [3]string{"bash", "-c", `hostname`}

	command := exec.Command(
		commandString[0],
		commandString[1],
		commandString[2],
	)
	rawOutput, _ := command.CombinedOutput()
	rawOutputString := strings.TrimSuffix(string(rawOutput), "\n")

	return rawOutputString
}

// Sends system information
func sendSystemInfoMacOS() SystemInfo {
	sysInfo := SystemInfo{}

	// Sets system type (i.e. operating system)
	sysInfo.SystemType = "MacOS"
	// Gets battery life %
	sysInfo.BatteryLifePercent = getBatteryPercentageMacOS()
	// Gets System Id
	sysInfo.SystemUser = getSystemUserMacOS()
	// Gets System Hostname
	sysInfo.SystemHostname = getSystemHostnameMacOS()

	return sysInfo
}

func sendSystemInfo(operatingSystem string) SystemInfo {
	var sysInfo SystemInfo

	// Sends system info to Master Node
	if operatingSystem == "MacOS" {
		sysInfo = sendSystemInfoMacOS()
	} else if operatingSystem == "Windows" {
		// TODO Add Windows functionality
	} else if operatingSystem == "Linux" {
		// TODO Add Linux functionality
	}

	return sysInfo
}

func getOperatingSystem() string {
	var operatingSystem string

	switch runtime.GOOS {
	case "darwin":
		operatingSystem = "MacOS"
	case "windows":
		operatingSystem = "Windows"
	case "linux":
		operatingSystem = "Linux"
	default:
		panic("ERROR: Couldn't detect Operating System!")
	}

	return operatingSystem
}

// Sends a polling message to Master Node
func sendPollingMessage(conn net.Conn) {
	sendMessage(conn, true, false, SystemInfo{})
}

// Creates and sends message to Master Node
func sendMessage(conn net.Conn, isPollMessage bool, isDataMessage bool, sysInfo SystemInfo) {
	message := SlaveNodeMessage{
		IsPollMessage: isPollMessage,
		IsDataMessage: isDataMessage,
		SystemInfo:    sysInfo,
		Timestamp:     time.Now().Unix(),
	}
	messageBytes, _ := json.Marshal(message)
	messageJson := string(messageBytes) + "\n"
	fmt.Fprintf(conn, messageJson)
}

// Converts Master Node message from JSON -> struct
func getMasterNodeMessageStruct(messageJson string) MasterNodeMessage {
	var messageStruct MasterNodeMessage
	json.Unmarshal([]byte(messageJson), &messageStruct)

	return messageStruct
}

func main() {
	const googleComputeIp = "35.243.155.9"
	const localIp = "127.0.0.1"
	const port = 3141

	//operatingSystem := getOperatingSystem()

	// Sets initial values
	pollingDelaySeconds := 5
	masterNodeResponse := MasterNodeMessage{
		WantResponse:false,
		WantSystemInfo:false,
		TimeDelayOnResponseSeconds:5,
		Timestamp:time.Now().Unix(),
	}
	for {
		// Connects to Master Node
		conn, err := net.Dial("tcp", localIp+":"+strconv.Itoa(port))
		if err != nil {
			panic(err)
		}

		// If Master Node doesn't want response, send polling message
		// Else, respond with the message the Master Node requested
		if masterNodeResponse.WantResponse==false{
			sendPollingMessage(conn)
		} else {
			systemInfo := SystemInfo{}

			// If System Info is requested, gets system information
			if masterNodeResponse.WantSystemInfo == true{
				// TODO: Add system info here
				systemInfo.SystemHostname = "test name please ignore"
			}

			// Sends message
			sendMessage(conn, false, true, systemInfo)
		}

		//sysInfo := sendSystemInfo(operatingSystem)
		//message := SlaveNodeMessage{
		//	IsPollMessage: false,
		//	IsDataMessage: true,
		//	SystemInfo:  sysInfo,
		//}
		//// Creates SystemInfo struct, then converts it to JSON
		//messageBytes, _ := json.Marshal(message)
		//messageJson := string(messageBytes) + "\n"
		//
		//// Sends JSON containing system info to Master Node
		//fmt.Fprintf(conn, messageJson)

		masterNodeResponseJson, _ := bufio.NewReader(conn).ReadString('\n')
		masterNodeResponse = getMasterNodeMessageStruct(masterNodeResponseJson)
		fmt.Print(masterNodeResponseJson)
		//fmt.Println(masterNodeResponse)
		fmt.Println(masterNodeResponse.WantResponse)

		// Sleeps for time requested by user
		pollingDelaySeconds = masterNodeResponse.TimeDelayOnResponseSeconds
		time.Sleep(time.Duration(pollingDelaySeconds) * time.Second)
	}
}

//
// TODO Add polling
// TODO Add support for WantTimeDelayOnResponse (MasterNodeMessage response)
