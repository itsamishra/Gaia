package main

import (
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

// Updates subNodeSlice with SubNode
func updateSubNodeStatus(w http.ResponseWriter, r *http.Request) {
	// Gets parameters from url
	ip := r.URL.Query().Get("IP")
	batteryLevelPercentage, err := strconv.ParseFloat(r.URL.Query().Get("BatteryLevelPercentage"), 64)
	handleError(err)

	// Creates SubNode struct
	subNode := SubNode{
		IP:                     ip,
		BatteryLevelPercentage: batteryLevelPercentage,
		UnixTimestamp:          time.Now().Unix(),
	}

	// Adds SubNode to subNodeSlice
	subNodeSlice[ip] = subNode
	fmt.Println("subNodeSlice:")
	fmt.Println(subNodeSlice)
	fmt.Printf("Number of SubNodes in subNodeSlice: %d\n", len(subNodeSlice))
	fmt.Println("-----------------------------------------------------------------")

	fmt.Fprintf(w, "SubNode Updated!")
}

type myData struct {
	Name  string
	Photo string
}

// Accepts POST from Sub Nodes and updates their details
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ping Handled")

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
	fmt.Println(subNode)

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
	const masterNodePort = "3141"

	// Updates subNodeSlice with SubNode
	http.HandleFunc("/api/update-sub-node", updateSubNodeStatus)
	http.HandleFunc("/api/ping", pingHandler)
	log.Fatal(http.ListenAndServe(":"+masterNodePort, nil))
}
