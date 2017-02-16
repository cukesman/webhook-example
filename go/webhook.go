package main

import (
	"log"
	"net/http"
)

type EventPayload struct {
	Identifier string `json:"id"`

	// One Off Execution
	Gherkin string `json:"gherkin"`

	// Version
	VersionNumber string `json:"versionNumber"`

	// Feature
	Title            string `json:"title"`
	FileName         string `json:"fileName"`
	PreviousFileName string `json:"previousFileName"`
}

type User struct {
	FirstName string       `json:"firstName"`
	LastName  string       `json:"lastName"`
	Email     string       `json:"email"`
	Payload   EventPayload `json:"payload"`
}

type Event struct {
	Type string `json:"event"`
	Date string `json:"date"`
	User User   `json:"user"`
}

var hookHandler = func(rw http.ResponseWriter, r *http.Request) {
	log.Println("Received call from " + r.Host + r.Method)
}

func main() {
	http.HandleFunc("/", hookHandler)

	http.ListenAndServe(":3600", nil)
}
