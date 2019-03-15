package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// Struct containing all important system information
type SystemInfo struct {
	SystemId              string
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
	// Port on which Master Node listens
	const port = 3141

	// Starts Master Node server
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		handleError(err)
	}
	fmt.Println("Listening on port", port)

	// Continuously listens for connection
	for {
		conn, err := ln.Accept()
		if err != nil {
			handleError(err)
		}

		// Gets message sent by Slave Node
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			handleError(err)
		}
		message = strings.TrimSuffix(message, "\n")

		var messageStruct SystemInfo
		json.Unmarshal([]byte(message), &messageStruct)
		fmt.Println("Message:", message)
		fmt.Println("Message Struct:", messageStruct)
	}
}
