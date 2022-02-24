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
	urlSplit := strings.Split(r.URL.Path, "/")
	var urlWant int
	comp := strings.ReplaceAll(NEIGHBOURUNIS_PATH, "/", "")
	//fmt.Println(urlSplit, ":", urlSplit[len(urlSplit)-1], ":", len(urlSplit))
	for i, s := range urlSplit {
		if s == (comp) {
			urlWant = i + 1
		}
	}
	var firstAppendVal, secondAppendVal string
	firstAppendVal = strings.ReplaceAll(urlSplit[urlWant], " ", "%20")
	if len(firstAppendVal) > 0 {
		firstAppendVal = "=" + firstAppendVal
	}
	if len(urlSplit)-1 == urlWant+1 {
		secondAppendVal = strings.ReplaceAll(urlSplit[urlWant+1], " ", "%20")
		if len(secondAppendVal) > 0 {
			secondAppendVal = "=" + secondAppendVal
		} else {
			secondAppendVal = ""
		}
	} else {
		secondAppendVal = ""
	}
	write, err := http.Get("http://universities.hipolabs.com/search?name" + secondAppendVal + "&country" + firstAppendVal)
	checkError(err)

	var getLimit int64
	getParam := strings.Split(r.URL.RawQuery, "limit=")
	if len(getParam) > 1 {
		t, _ := strconv.ParseInt(getParam[1], 10, 0)
		getLimit = t
	}

	if getLimit == 0 {
		//fmt.Println(getLimit) //if size = 0 set to default.
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
	setUni = append(setUni, setUniversity(getU)...)
	setUni = append(setUni, getBorderingUniversities(setUni, int(getLimit))...)

	// Encode specific content --> Alternative: "err := json.NewEncoder(w).Encode(location)"
	err = encoder.Encode(setUni)
	checkError(err)

}
