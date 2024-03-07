package Handles

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// StatusHandler manages requests to the librarystats/v1/status/ endpoint
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"gutendexapi":  getStatusCode("http://129.241.150.113:8000/books"),
		"languageapi":  getStatusCode("http://129.241.150.113:3000/"),
		"countriesapi": getStatusCode("http://129.241.150.113:8080/v3.1/all/"),
		"version":      "v1",
	}

	// Converts the status into JSON format using marshaling
	jsonRes, err := json.Marshal(status)
	if err != nil {
		http.Error(w, "Unable to marshal JSON response", http.StatusInternalServerError)
		return
	}

	// Configures the content type header and sends the likely response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func getStatusCode(url string) int {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return http.StatusInternalServerError
	}
	defer response.Body.Close()
	return response.StatusCode
}
