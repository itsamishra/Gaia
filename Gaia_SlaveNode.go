package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type systemInfo struct {
	SystemId string `json:"empname"`
	SystemType string
	BatteryLifePercent float64
	BatteryHoursRemaining float64
}

func handleError(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3141")
	if err != nil {
		handleError(err)
	}

	systemInfo := &systemInfo{
		SystemId:"SYSTEM001",
		SystemType:"MacOS",
		BatteryLifePercent:64,
		BatteryHoursRemaining:2.5,
	}
	jsonBytes, _ := json.Marshal(systemInfo)
	var jsonStr = string(jsonBytes) + "\n"
	fmt.Printf(jsonStr)

	//fmt.Fprintf(conn, "This is a sample message\n")
	fmt.Fprintf(conn, jsonStr)

}

// TODO figure out how to send JSON
// TODO send sample JSON
// TODO send JSON containing macbook username, battery level, battery time remaining, etc.
