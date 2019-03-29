package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getOpenPorts() []string {
	// Bash file containing command that gets battery level
	cmd := "./Bash Functions/getOpenPorts.sh"

	// Gets battery level
	openPortsByte, _ := exec.Command(cmd).Output()
	openPortsString := string(openPortsByte)
	openPortsString = strings.Trim(openPortsString, "\n")
	openPortsArray := strings.Split(openPortsString, "\n")

	return openPortsArray
}

func main() {
	fmt.Println(getOpenPorts())
}
