package Handles

import (
	"assignment-1/Structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// HandlerBookCount manages the requests to the /librarystats/v1/bookcount/ endpoint
func HandlerBookCount(w http.ResponseWriter, r *http.Request) {
	language := r.URL.Query().Get("language")
	if language == "" {
		http.Error(w, "A language is mandatory!", http.StatusBadRequest)
		return
	}

	// Invoke a function to obtain the book count for the specific language(s)
	bkCount, err := getBookCount(language)
	if err != nil {
		http.Error(w, "Unable to obtain book count information!", http.StatusInternalServerError)
		return
	}

	jsonRes, err := json.Marshal(bkCount)
	if err != nil {
		http.Error(w, "Unable to marshal JSON response!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

// Function that fetches the book count for your specific language(s)
func getBookCount(language string) (interface{}, error) {
	var bookCounts []interface{}

	// Loop to that iterates through each language in the language list of the Gutendex endpoint
	for _, lang := range strings.Split(language, ",") {
		uniqAuthors := make(map[string]bool)

		response, err := http.Get("http://129.241.150.113:8000/books/?languages=" + lang)
		if err != nil {
			fmt.Print("Error")
		}
		var book Structs.Book
		err = json.NewDecoder(response.Body).Decode(&book)
		if err != nil {
			fmt.Println("Error:", err)
		}

		allBooks, err := http.Get("http://129.241.150.113:8000/books")
		if err != nil {
			fmt.Print("Error")
		}
		var totalBooks Structs.Book
		err = json.NewDecoder(allBooks.Body).Decode(&totalBooks)
		if err != nil {
			fmt.Println("Error:", err)
		}
		gutenlibBooks := totalBooks.Count

		fmt.Print(book.Count)
		pageNum := book.Count / 32
		if book.Count%32 != 0 {
			pageNum += 1
		}

		// Traverses through all the pages in the Gutendex endpoint and adds it to author map
		for page := 1; page <= pageNum; page++ {
			response, err := http.Get("http://129.241.150.113:8000/books/?languages=" + lang + "&page=" + strconv.Itoa(page))
			if err != nil {
				fmt.Print("Error")
			}
			var book Structs.Book
			err = json.NewDecoder(response.Body).Decode(&book)
			if err != nil {
				fmt.Println("Error:", err)
			}
			for _, result := range book.Results {
				for _, author := range result.Authors {
					fmt.Print(author.Name)
					if !uniqAuthors[author.Name] {
						uniqAuthors[author.Name] = true
					}
				}
			}
		}
		// Outputs the total number of authors
		fmt.Print(len(uniqAuthors))

		// Sets up a map to store book count information for the selected language(s)
		bkCount := map[string]interface{}{
			"language": lang,             // Language(s) requested by you
			"books":    book.Count,       // Total number of books in the specified language
			"authors":  len(uniqAuthors), // All unique authors in that specific language
			"fraction": float32(book.Count) / float32(gutenlibBooks),
		}

		// Adds the book count details to the bookCounts slice
		bookCounts = append(bookCounts, bkCount)
	}

	// Provides the bookCounts slice, which holds book count data for all languages
	return bookCounts, nil
}
