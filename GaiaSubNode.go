package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

// Sub Node message
type SubNodeMessage struct {
	MessageType string // Can be "Poll" or "Data"
}

func main() {
	// Constants
	const googleComputeIp = "35.243.155.9"
	const localIp = "127.0.0.1"
	const port = 3141

	// Minimum wait time between calls to Master Node
	// pollingDelaySeconds := 0.5

	for {
		// fmt.Println("loop")

		// Connects to Master Node
		conn, err := net.Dial("tcp", localIp+":"+strconv.Itoa(port))
		handleError(err)

		message, err := bufio.NewReader(conn).ReadString('\n')
		handleError(err)
		fmt.Print(message)

		// // Sends message to Master Node
		// _, err = fmt.Fprintf(conn, "This is a test message\n")
		// handleError(err)

		// // Reads response from Master Node
		// message, err := bufio.NewReader(conn).ReadString('\n')
		// handleError(err)
		// fmt.Print(message)

		// Sleeps for set number of seconds
		// time.Sleep(time.Duration(1000*pollingDelaySeconds) * time.Millisecond)
	}
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
