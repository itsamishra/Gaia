package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type NodeData struct {
	IP           string
	BatteryLevel float64
}

// Adds header to prevent Access-Control-Allow-Origin CORS error
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getNodeStatus(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	fmt.Println(params)
	queryParams := r.URL.Query()
	fmt.Println(queryParams)
	fmt.Println(queryParams.Get("allNodes"))

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

func main() {
	const port = 3141

	router := mux.NewRouter()

	// Look over https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo to set POST/etc.
	router.HandleFunc("/api/node-status", getNodeStatus).Methods("GET")
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
