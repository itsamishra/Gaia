package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	//"strings"
)

type MasterNodeMessage struct {
	WantResponse               bool
	WantSystemInfo             bool
	WantTimeDelayOnResponse    bool
	TimeDelayOnResponseSeconds int
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
		//fmt.Printf(message)

		var messageStruct SystemInfo
		json.Unmarshal([]byte(message), &messageStruct)

		// Asks for response
		masterNodeResponse := MasterNodeMessage{
			WantResponse:               true,
			WantSystemInfo:             true,
			WantTimeDelayOnResponse:    false,
			TimeDelayOnResponseSeconds: 0,
		}
		jsonBytes, _ := json.Marshal(masterNodeResponse)
		var jsonStr = string(jsonBytes) + "\n"

		conn.Write([]byte(jsonStr))

	}
}
