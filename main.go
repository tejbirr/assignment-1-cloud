package main

import (
	"assignment-1/Handles"
	"net/http"
)

func main() {

	http.HandleFunc("/librarystats/v1/bookcount/", Handles.HandlerBookCount)
	http.HandleFunc("/librarystats/v1/status/", Handles.StatusHandler)
	http.HandleFunc("/librarystats/v1/readership/", Handles.HandlerReadership)
	http.ListenAndServe(":8080", nil)

}
