package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Structs
type NodeData struct {
	IP           string
	BatteryLevel float64
}
type SubNode struct {
	Url           string
	Name          string
	IP            string
	UnixTimeAdded int64
}

// Adds header to prevent Access-Control-Allow-Origin CORS error
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Returns Sub Node data
func getNodeStatus(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	queryParams := r.URL.Query()
	allNodes, err := strconv.ParseBool(queryParams.Get("allNodes"))
	handleError(err)

	fmt.Println(allNodes)

	nodeData := NodeData{
		IP:           "123.456.789.0",
		BatteryLevel: 76.5,
	}

	nodeDataSlice := make([]NodeData, 0)
	nodeDataSlice = append(nodeDataSlice, nodeData)
	nodeDataSlice = append(nodeDataSlice, nodeData)
	nodeDataSlice = append(nodeDataSlice, nodeData)

	json.NewEncoder(w).Encode(nodeDataSlice)
}

// Adds Sub Node to list of Sub Nodes (with timestamp)
func addSubNodeToList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// Uses query params to create Sub Node
	queryParams := r.URL.Query()
	subNode := SubNode{
		Url:           queryParams.Get("url"),
		Name:          queryParams.Get("name"),
		IP:            queryParams.Get("ip"),
		UnixTimeAdded: time.Now().Unix(),
	}

	// Adds a new Sub Node to subNodesConnected
	subNodesConnected[queryParams.Get("url")] = subNode
	// fmt.Println(subNodesConnected)
	// fmt.Println(len(subNodesConnected))
}

var subNodesConnected = make(map[string]SubNode)

func main() {
	const port = 3141

	// Sets up /api/node-status endpoint
	router := mux.NewRouter()
	// Look over https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo to set POST/etc.
	router.HandleFunc("/api/node-status", getNodeStatus).Methods("GET")
	router.HandleFunc("/api/check-in", addSubNodeToList).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

// If error is passed in, throws error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO: Set up connection between Node server & Golang server
// TODO: Write a function that listens for new Sub Nodes on port 3142 (note that a list of Sub Nodes must be maintained on the Master Node)
// TODO: Write a function that communicates with the Sub Node on port 3143
