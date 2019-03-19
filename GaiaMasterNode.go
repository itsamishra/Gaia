package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type Job struct {
	IP string
}

func handleSubNodeConnection(conn net.Conn) {
	// Sends message to Sub Node
	_, err := fmt.Fprintf(conn, "This is a message from the Master Node\n")
	handleError(err)
}

// Listens for new jobs and adds them to jobSlice
func listenForNewJobs() {

}

var jobSlice = make([]Job, 0)

func main() {
	const port = 3141

	fmt.Println(jobSlice)

	// Creates listener
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	handleError(err)

	// Connects to
	for {
		// Creates connection to Sub Node
		conn, err := ln.Accept()
		handleError(err)

		// Artificial pause
		time.Sleep(500 * time.Millisecond)

		go handleSubNodeConnection(conn)
	}
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO: See new picture --> main for loop should be listening for request from Express app
