package main

import (
	"log"
	"net/http"
	"os"
	"restfultest/handler"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href="` + (handler.RESOURCE_ROOT_PATH + handler.NEIGHBOURUNIS_PATH) + `">Neighbouring universities</a><br>`))
	w.Write([]byte(`<a href="` + (handler.RESOURCE_ROOT_PATH + handler.UNIINFO_PATH) + `">University info</a><br>`))
	w.Write([]byte(`<a href="` + (handler.RESOURCE_ROOT_PATH + handler.DIAG_PATH) + `">Diagnostics interface</a><br>`))

}

func handleRequests(port string) {

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func setPort(inn string) {
	os.Setenv("PORT", inn)
}

func getPort() string {
	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}
	return port
}

func main() {
	setPort("") // default port is :8080

	// Set up handler endpoints
	http.HandleFunc(handler.RESOURCE_ROOT_PATH, homePage)
	http.HandleFunc(handler.DEFAULT_PATH, handler.EmptyHandler)
	http.HandleFunc(handler.RESOURCE_ROOT_PATH+handler.NEIGHBOURUNIS_PATH, handler.LocationHandler)
	http.HandleFunc(handler.RESOURCE_ROOT_PATH+handler.UNIINFO_PATH, handler.CollectionHandler)
	http.HandleFunc(handler.RESOURCE_ROOT_PATH+handler.DIAG_PATH, handler.CollectionHandler)

	handleRequests(getPort())
}
