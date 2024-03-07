# Assignment-1
Hi,
    This is my first assignment for the cloud technologies course. In this assignment we were asked to develop a REST web application in Go that that gives our users the details about books from Gutenberg Library (in a specific language). Additionally, the application should estimate the number of potential readers for that language using external web services, and present the results in a specified format.

## Datiled Description
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
- This assignment was challenging for me, but I found it enjoyable overall. Initially, I underestimated it, thinking it was the first one and took it lightly. However, as I started working on it, I realized it required more time than I expected. Moving forward, I'll be more careful and allocate enough time for upcoming assignments.
