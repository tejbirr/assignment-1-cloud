# Cloud Technologies - Assignment 1: REST Web Application in Go
This project is my first assignment for the Cloud Technologies course. The task was to develop a RESTful web application in Go that retrieves book details from the Gutenberg Library for a specific language. Additionally, the application estimates the number of potential readers for that language using external web services and presents the results in a structured format.

## Detailed Description
A list of features:
My web service have three resource root paths
- /librarystats/v1/bookcount/
- /librarystats/v1/readership/
- /librarystats/v1/status/

This service will run on localhost port 8080, so my resource root paths look like this
- http://localhost:8080/librarystats/v1/bookcount/
- http://localhost:8080/librarystats/v1/readership/
- http://localhost:8080/librarystats/v1/status/

My Bookcount Endpoint:
- The first endpoint counts and returns the number of books and unique authors for a specified language, which can be one or multiple languages (identified by 2-letter ISO codes, separated by commas).

Example request: bookcount/?language=no,fi

Body (Example): 
[
  {
     "language": "no",
     "books": 21,
     "authors": 14,
     "fraction": 0.0005
  }
]

My Readership Endpoint:
- The second endpoint needs to tell you how many people could potentially read books in a particular language. It does this by looking at the population of countries where the language is official. The result should include the number of books, unique authors, and this estimated reader count.

Example request: readership/no

Body (Example): 
[ 
  {
     "country": "Norway",
     "isocode": "NO",
     "books": 21,
     "authors": 14,
     "readership": 5379475
  }
]

My Status Endpoint:
- The diagnostics interface shows whether the services that our service relies on are available or not. It does this by checking the status codes returned by those services. Additionally, it should give us information about how long our service has been running, known as uptime. I faced some challenges with this assignment, and as a result, this part is still incomplete. However, I have a strong interest in the overall subject and really looking forward to learning more.

Body Example:
{
   "gutendexapi": "<http status code for gutendex API>",
   "languageapi: "<http status code for language2countries API>", 
   "countriesapi": "<http status code for restcountries API>",
   "version": "v1",
}

## General Thought:
- This assignment was more challenging than I initially expected. I underestimated its complexity, assuming it would be straightforward since it was the first one. However, as I progressed, I realized it required more time and effort than anticipated. This experience has taught me to approach future assignments with better time management and careful planning.
