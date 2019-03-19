package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	const port = 3141

	// Creates listener
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	handleError(err)

	// Infinite loop - listens for connections forever
	for {
		// Creates connection to Sub Node
		conn, err := ln.Accept()
		handleError(err)
		// fmt.Println("test")

		time.Sleep(5 * time.Second)

		// Sends message
		fmt.Fprintf(conn, "This is a message from the Master Node\n")

		// // Reads incoming message
		// message, err := bufio.NewReader(conn).ReadString('\n')
		// handleError(err)
		// fmt.Print(message)

		// // Responds to Sub Node
		// uMessage := strings.ToUpper(message)
		// fmt.Fprintf(conn, uMessage)
	}
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
