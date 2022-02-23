package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
Entry point handler for Location information
*/
func UniinfoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		//handlePostRequest(w, r)
	case http.MethodGet:
		handleGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET or POST are supported.", http.StatusNotImplemented)
		return
	}

}

/*
Dedicated handler for POST requests
*/
func handlePostRequest(w http.ResponseWriter, r *http.Request) {

	// Instantiate decoder
	decoder := json.NewDecoder(r.Body)
	universities := Universities{}

	// Decode location instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&location)"
	err := decoder.Decode(&universities)
	if err != nil {
		http.Error(w, "Error during decoding", http.StatusInternalServerError)
		return
	}

	// Flat printing
	fmt.Println("Received following request:")
	fmt.Println(universities)

	// Pretty printing
	output, err := json.MarshalIndent(universities, "", "  ")
	if err != nil {
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return
	}

	fmt.Println("Pretty printing:")
	fmt.Println(string(output))

	// TODO: Handle content (e.g., writing to DB, process, etc.)

	// Return status code (good practice)
	http.Error(w, "OK", http.StatusOK)
}

/*
Dedicated handler for GET requests
*/
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	urlSplit := strings.Split(r.URL.Path, "/")
	var urlWant int
	//fmt.Println(urlSplit, ":", urlSplit[len(urlSplit)-1], ":", len(urlSplit))
	for i, s := range urlSplit {
		if s == "uniinfo" {
			urlWant = i + 1
		}
	}

	lastAppendVal := strings.ReplaceAll(urlSplit[urlWant], " ", "%20")
	write, err := http.Get("http://universities.hipolabs.com/search?name=" + lastAppendVal)
	checkError(err)
	var getU []getUnii
	body, err := io.ReadAll(write.Body)

	checkError(err)
	json.Unmarshal(body, &getU)
	// Write content type header (best practice)
	w.Header().Add("content-type", "application/json")

	// Instantiate encoder
	encoder := json.NewEncoder(w)

	// Encode specific content --> Alternative: "err := json.NewEncoder(w).Encode(location)"
	err = encoder.Encode(setUniversity(getU))
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}
}
