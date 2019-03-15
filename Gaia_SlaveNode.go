package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

// Struct containing all important system information
type SystemInfo struct {
	SystemId              string `json:"empname"`
	SystemType            string
	BatteryLifePercent    float64
	BatteryHoursRemaining float64
}

// Sends system information
func sendSystemInfo(conn net.Conn) {
	// Creates SystemInfo struct, then converts it to JSON
	si := SystemInfo{
		SystemId:              "SYSTEM001",
		SystemType:            "MacOS",
		BatteryLifePercent:    64,
		BatteryHoursRemaining: 2.5,
	}
	jsonBytes, _ := json.Marshal(si)
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

func main() {
	const googleComputeIp = "35.243.155.9"
	const localIp = "127.0.0.1"
	const port = 3141

	// Connects to Master Node
	conn, err := net.Dial("tcp", localIp+":"+strconv.Itoa(port))
	if err != nil {
		handleError(err)
	}

	// Sends system info to Master Node
	sendSystemInfo(conn)
}

// TODO send JSON containing macbook username, battery level, battery time remaining, etc.
