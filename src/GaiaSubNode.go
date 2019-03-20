package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	// Constants
	const googleComputeIp = "35.243.155.9"
	const localIp = "127.0.0.1"
	const port = 3144

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	handleError(err)
	for {
		conn, err := ln.Accept()
		handleError(err)

		fmt.Println(conn)

	}
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
