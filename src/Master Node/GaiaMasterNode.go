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
	IP                     string
	BatteryLevelPercentage float64
	Screenshot             string // String??? Or some other datatype???
	UnixTimestamp          int64
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

func ping(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
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
	http.HandleFunc("/api/ping", ping)
	log.Fatal(http.ListenAndServe(":"+masterNodePort, nil))
}
