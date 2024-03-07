package Handles

import (
	"assignment-1/Structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// HandlerReadership manages requests to the /librarystats/v1/readership/ endpoint
func HandlerReadership(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	var Dataarray []Structs.ReadershipData
	parts := strings.Split(url, "/")

	// Retrieves information for a country using its country code
	langCode := parts[len(parts)-1]
	fmt.Println(langCode)
	response, err := http.Get("http://129.241.150.113:3000/language2countries/" + langCode)
	if err != nil {
		fmt.Print("Error")

	}
	var country []Structs.Country
	err = json.NewDecoder(response.Body).Decode(&country)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer response.Body.Close()

	fmt.Println(country[0].Official_Name)

	// Uses the local endpoint (localhost:8080) in my program to fetch data for a particular language
	response, err = http.Get("http://localhost:8080/librarystats/v1/bookcount?language=" + langCode)
	if err != nil {
		fmt.Println("Error")
	}
	var bookCounts []Structs.LanguageData
	err = json.NewDecoder(response.Body).Decode(&bookCounts)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer response.Body.Close()

	// Loop iterating through each country to obtain the required data
	for i := 0; i < len(country); i++ {
		response, err = http.Get("http://129.241.150.113:8080/v3.1/alpha/" + country[i].ISO3166Alpha2)
		if err != nil {
			fmt.Print("Error")

		}
		var countries []Structs.Population
		err = json.NewDecoder(response.Body).Decode(&countries)
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer response.Body.Close()

		// Outputs the relevant data for each country with the specified country code until the loop finishes
		fmt.Println("\n\n" + country[i].ISO3166Alpha2)
		fmt.Println(country[i].Official_Name)
		fmt.Println(bookCounts[0].Books)
		fmt.Println(bookCounts[0].Authors)
		fmt.Println(countries[0].Population)

		output := Structs.ReadershipData{country[i].ISO3166Alpha2,
			country[i].Official_Name,
			bookCounts[0].Books,
			bookCounts[0].Authors,
			countries[0].Population}

		Dataarray = append(Dataarray, output)
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(Dataarray); err != nil {
		fmt.Fprintf(w, "Error encoding json")
	}
}
