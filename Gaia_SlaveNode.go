package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// Struct containing all important system information
type SystemInfo struct {
	SystemUser         string
	SystemType         string
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

// Sends system information
func sendSystemInfoMacOS(conn net.Conn) {
	sysInfo := SystemInfo{}

	// Sets system type (i.e. operating system)
	sysInfo.SystemType = "MacOS"
	// Gets battery life %
	sysInfo.BatteryLifePercent = getBatteryPercentageMacOS()
	// Gets System Id
	sysInfo.SystemUser = getSystemUserMacOS()

	// Creates SystemInfo struct, then converts it to JSON
	jsonBytes, _ := json.Marshal(sysInfo)
	var jsonStr = string(jsonBytes) + "\n"
	fmt.Printf(jsonStr)

	// Sends JSON containing system info to Master Node
	fmt.Fprintf(conn, jsonStr)
}

// Prints error, then exits with Exit Code 1
func handleError(err error) {
	log.Fatal(err)
	os.Exit(1)
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

func main() {
	const googleComputeIp = "35.243.155.9"
	const localIp = "127.0.0.1"
	const port = 3141

	operatingSystem := getOperatingSystem()
	fmt.Println(operatingSystem)

	// Connects to Master Node
	conn, err := net.Dial("tcp", localIp+":"+strconv.Itoa(port))
	if err != nil {
		handleError(err)
	}

	// Sends system info to Master Node
	if operatingSystem == "MacOS" {
		sendSystemInfoMacOS(conn)
	}
}
