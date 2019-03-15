package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
	// Port on
	listenPort1 := "127.0.0.1:8081"
	//listenPort2 := "35.243.155.9:8081"
	// connect to this socket
	conn, _ := net.Dial("tcp", listenPort1)
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
