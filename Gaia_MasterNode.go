package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

// Message sent
type MasterNodeMessage struct {
	WantResponse               bool
	WantSystemInfo             bool
	TimeDelayOnResponseSeconds int
	Timestamp                  int64 // When was this sent?
}

// Message received
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

func serveStaticHtml() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(":3000", nil)
}

func main() {
	// Port on which Master Node listens
	const port = 3141

	// Serves static HTML page
	//go serveStaticHtml()

	// Starts Master Node server
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port", port)

	// Continuously listens for connection
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// Gets message sent by Slave Node
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print("Received Message:\n" + message)

		var messageStruct SystemInfo
		json.Unmarshal([]byte(message), &messageStruct)

		// Asks for response
		masterNodeResponse := MasterNodeMessage{
			WantResponse:               false,
			WantSystemInfo:             false,
			TimeDelayOnResponseSeconds: 3,
			Timestamp:                  time.Now().Unix(),
		}
		jsonBytes, _ := json.Marshal(masterNodeResponse)
		var jsonStr = string(jsonBytes) + "\n"

		conn.Write([]byte(jsonStr))

	}
}
