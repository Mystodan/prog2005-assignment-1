package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

/*
Entry point handler for Location information
*/
func NBuinfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		NBGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET or POST are supported.", http.StatusNotImplemented)
		return
	}
}

/*
Dedicated handler for GET requests
*/
func NBGetRequest(w http.ResponseWriter, r *http.Request) {
	urlSplit := strings.Split(r.URL.Path, "/") // splits url into readable strings
	var urlWant int
	errCode := false //sets for error
	comp := strings.ReplaceAll(NEIGHBOURUNIS_PATH, "/", "")
	for i, s := range urlSplit {
		if s == (comp) {
			urlWant = i + 1
		}
	}
	var firstAppendVal, secondAppendVal string
	firstAppendVal = strings.ReplaceAll(urlSplit[urlWant], " ", "%20") // for replacing spaces with ascii
	if len(firstAppendVal) > 0 {                                       // if the length of country input is larger than 0
		firstAppendVal = "=" + firstAppendVal
	} else {
		errCode = true
	}
	if len(urlSplit)-1 == urlWant+1 { // if the url size is wanted
		secondAppendVal = strings.ReplaceAll(urlSplit[urlWant+1], " ", "%20")
		if !(len(secondAppendVal) > 0) { // if the length of name input is not wanted
			errCode = true
		}
	} else {
		errCode = true
	}

	if errCode { // err implementation
		http.Error(w, "No functionality without parameters: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}", http.StatusOK)
	} else {

		write := getURL(GET_UNI + UNI_REQ + secondAppendVal + "&country" + firstAppendVal) // gets value from api

		var getLimit int64
		getParam := strings.Split(r.URL.RawQuery, "limit=") // query for optional parameter
		if len(getParam) > 1 {                              // if parameter there is a parameter
			t, _ := strconv.ParseInt(getParam[1], 10, 0) // convert to int
			getLimit = t
		}

		var getU []getUnii
		body, err := io.ReadAll(write.Body)

		checkError(err)
		json.Unmarshal(body, &getU)
		// Write content type header (best practice)
		w.Header().Add("content-type", "application/json")

		// Instantiate encoder
		encoder := json.NewEncoder(w)
		var setUni []Universities
		setUni = append(setUni, setUniversity(getU)...)                             //append universities
		setUni = append(setUni, getBorderingUniversities(setUni, int(getLimit))...) //append bordering universities

		// Encode specific content --> Alternative: "err := json.NewEncoder(w).Encode(location)"
		err = encoder.Encode(setUni)
		checkError(err)
	}
}
