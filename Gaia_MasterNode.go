package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func handleError(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func main() {
	const port = ":3141"

	// Starts server
	ln, err := net.Listen("tcp", port)
	if err != nil {
		handleError(err)
	}
	fmt.Println("Listening on port", port)

	// Listens for connection
	for {
		conn, err := ln.Accept()
		if err != nil {
			handleError(err)
		}

		// Gets message sent by client
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			handleError(err)
		}
		message = strings.TrimSuffix(message, "\n")

		fmt.Println("Message:", message)
	}
}
