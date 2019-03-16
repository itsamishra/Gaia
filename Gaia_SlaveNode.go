package main

import (
	"bufio"
	"encoding/json"

	//"encoding/json"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Meessage sent
type SlaveNodeMessage struct {
	IsPollMessage bool // Is this message polling the server?
	IsDataMessage bool // Is this message returning data?
	SystemInfo    SystemInfo
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
func sendPollingMessage(conn net.Conn){
	sendMessage(conn, true, false, SystemInfo{})
}

// Creates and sends message to Master Node
func sendMessage(conn net.Conn, isPollMessage bool, isDataMessage bool, sysInfo SystemInfo){
	message := SlaveNodeMessage{
		IsPollMessage:isPollMessage,
		IsDataMessage:isDataMessage,
		SystemInfo:sysInfo,
	}
	messageBytes, _ := json.Marshal(message)
	messageJson := string(messageBytes) + "\n"
	fmt.Fprintf(conn, messageJson)
}

func main() {
	const googleComputeIp = "35	.243.155.9"
	const localIp = "127.0.0.1"
	const port = 3141
	const pollingDelaySeconds = 5

	//operatingSystem := getOperatingSystem()

	for {
		// Connects to Master Node
		conn, err := net.Dial("tcp", localIp+":"+strconv.Itoa(port))
		if err != nil {
			panic(err)
		}

		// Sends poll message
		sendPollingMessage(conn)


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
		fmt.Print(masterNodeResponseJson)

		time.Sleep(time.Duration(pollingDelaySeconds) * time.Second)
	}
}


// TODO Add polling
// TODO Add support for WantTimeDelayOnResponse (MasterNodeMessage response)