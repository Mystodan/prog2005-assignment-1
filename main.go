package main

import (
	"assignment-1/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}
	handler.TimerStart()
	// Set up handler endpoints
	http.HandleFunc(handler.DEFAULT_PATH, handler.EmptyHandler)
	http.HandleFunc(handler.RESOURCE_ROOT_PATH+handler.UNIINFO_PATH, handler.UniinfoHandler)
	http.HandleFunc(handler.RESOURCE_ROOT_PATH+handler.NEIGHBOURUNIS_PATH, handler.NBuinfoHandler)
	http.HandleFunc(handler.RESOURCE_ROOT_PATH+handler.DIAG_PATH, handler.DiagHandler)

	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
