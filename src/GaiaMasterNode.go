package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type SubNode struct {
	IP                     string
	BatteryLevelPercentage float64
}

// Updates subNodeSlice with SubNode
func updateSubNodeStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Params:")
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL.Query().Get("BatteryLevelPercentage"))

	// Creates SubNode struct and adds it to subNodeSlice
	batteryLevelPercentage, err := strconv.ParseFloat(r.URL.Query().Get("BatteryLevelPercentage"), 64)
	handleError(err)
	subNode := SubNode{
		BatteryLevelPercentage: batteryLevelPercentage,
	}
	fmt.Println(subNode)

	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "SubNode Updated!")
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Stores map of IP (string) to SubNode structs
var subNodeSlice = make(map[string]SubNode)

func main() {
	const port = "3141"

	// Updates subNodeSlice with SubNode
	http.HandleFunc("/api/update-sub-node", updateSubNodeStatus)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
