package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

func sendMessageToServer() {

}

func main() {
	const pollingDelaySeconds = 5

	for {
		conn, _ := net.Dial("tcp", "127.0.0.1:8081")

		//reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Text to send: ")
		//text, _ := reader.ReadString('\n')

		//fmt.Fprintf(conn, text+"\n")
		fmt.Fprintf(conn, strconv.Itoa(int(time.Now().Unix()))+"\n")
		str, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print(str)

		time.Sleep(time.Duration(pollingDelaySeconds) * time.Second)

	}
}
