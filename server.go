package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listening on 8181...")
	ln, _ := net.Listen("tcp", ":8081")

	for {
		conn, _ := ln.Accept()
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		fmt.Println([]byte(message + "\n"))

		conn.Write([]byte(message + "\n"))
	}
}
