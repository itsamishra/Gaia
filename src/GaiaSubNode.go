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

	for {
		// Connects to Master Node
		conn, err := net.Dial("tcp", localIp+":"+strconv.Itoa(port))
		handleError(err)

		// Gets message from Master Node
		message, err := bufio.NewReader(conn).ReadString('\n')
		handleError(err)
		fmt.Print(message)
	}
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
