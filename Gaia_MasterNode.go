package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// Prints error, then exits with Exit Code 1
func handleError(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func main() {
	// Port on which master node listens
	const port = ":3141"

	// Starts master node server
	ln, err := net.Listen("tcp", port)
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

		// Gets message sent by slave node
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			handleError(err)
		}
		message = strings.TrimSuffix(message, "\n")

		fmt.Println("Message:", message)
	}
}

// TODO Figure out how to parse JSON
