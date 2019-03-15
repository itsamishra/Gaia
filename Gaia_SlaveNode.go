package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

// Struct containing all important system information
type systemInfo struct {
	SystemId              string `json:"empname"`
	SystemType            string
	BatteryLifePercent    float64
	BatteryHoursRemaining float64
}

// Prints error, then exits with Exit Code 1
func handleError(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func main() {
	// Connects to master node
	conn, err := net.Dial("tcp", "127.0.0.1:3141")
	if err != nil {
		handleError(err)
	}

	// Creates systemInfo struct, then converts it to JSON
	systemInfo := systemInfo{
		SystemId:              "SYSTEM001",
		SystemType:            "MacOS",
		BatteryLifePercent:    64,
		BatteryHoursRemaining: 2.5,
	}
	jsonBytes, _ := json.Marshal(systemInfo)
	var jsonStr = string(jsonBytes) + "\n"
	fmt.Printf(jsonStr)

	// Sends JSON containing system info to master node
	fmt.Fprintf(conn, jsonStr)

}

// TODO send JSON containing macbook username, battery level, battery time remaining, etc.
