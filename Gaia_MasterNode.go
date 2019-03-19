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
	WantDataResponse           bool // Tells Slave Node if data is requested
	WantSystemInfo             bool // Tells Slave Node if System Info is requested
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

// Converts Slave Node message from JSON -> struct
func getSlaveNodeMessageStruct(messageJson string) SlaveNodeMessage {
	var messageStruct SlaveNodeMessage
	json.Unmarshal([]byte(messageJson), &messageStruct)

	return messageStruct
}

func serveStaticHtml() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(":3000", nil)
}

// Creates and sends message to Slave Node
func sendMessage(conn net.Conn, wantDataResponse bool, wantSystemInfo bool, timeDelayOnResponseSeconds int) {
	message := MasterNodeMessage{
		WantDataResponse:           wantDataResponse,
		WantSystemInfo:             wantSystemInfo,
		TimeDelayOnResponseSeconds: timeDelayOnResponseSeconds,
	}
	messageBytes, _ := json.Marshal(message)
	messageJson := string(messageBytes) + "\n"
	fmt.Fprintf(conn, messageJson)

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
		fmt.Println(getSlaveNodeMessageStruct(message))

		var messageStruct SystemInfo
		json.Unmarshal([]byte(message), &messageStruct)

		// Asks for response
		masterNodeResponse := MasterNodeMessage{
			WantDataResponse:           true,
			WantSystemInfo:             true,
			TimeDelayOnResponseSeconds: 3,
			Timestamp:                  time.Now().Unix(),
		}
		jsonBytes, _ := json.Marshal(masterNodeResponse)
		var jsonStr = string(jsonBytes) + "\n"

		conn.Write([]byte(jsonStr))

	}
}
