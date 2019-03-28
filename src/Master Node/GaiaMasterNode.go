package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// SubNode struct contains system information from machine associated with Sub Node
type SubNode struct {
	IP                      string
	BatteryLevelPercentage  float64
	Base64EncodedScreenshot string
	UnixTimestamp           int64
}

// Accepts POST from Sub Nodes and updates their details
func pingHandler(w http.ResponseWriter, r *http.Request) {
	// Parses parameters
	r.ParseForm()
	ip := r.Form.Get("IP")
	base64EncodedScreenshot := r.Form.Get("Base64Image")
	batteryLevelPercentage, err := strconv.ParseFloat(r.Form.Get("BatteryLevelPercentage"), 64)
	handleError(err)

	// Creates SubNode struct
	subNode := SubNode{
		IP:                      ip,
		BatteryLevelPercentage:  batteryLevelPercentage,
		Base64EncodedScreenshot: base64EncodedScreenshot,
		UnixTimestamp:           time.Now().Unix(),
	}

	// Adds SubNode to subNodeMap
	subNodeMap[ip] = subNode
	fmt.Println(subNodeMap)
	fmt.Printf("Number of SubNodes in subNodeMap: %d\n", len(subNodeMap))
	fmt.Println("-----------------------------------------------------------------")

	fmt.Fprintf(w, "SubNode Updated!")
}

func getInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Converts subNodeMap to JSON
	subNodeJSON, err := json.Marshal(subNodeMap)
	handleError(err)

	// Sends subNodeJSON
	fmt.Fprintf(w, string(subNodeJSON))
}

// Handles error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Stores map of IP (string) to SubNode structs
var subNodeMap = make(map[string]SubNode)

func main() {
	// Node which the Master Node listens on
	const masterNodePort = "3141"

	// Updates subNodeMap with SubNode
	http.HandleFunc("/api/ping", pingHandler)
	// Sends out Sub Node information to Dashboard
	http.HandleFunc("/api/getInfo", getInfoHandler)
	log.Fatal(http.ListenAndServe(":"+masterNodePort, nil))
}
