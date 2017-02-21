package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type eventPayload struct {
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

type user struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// Event is triggered by Cukesman
type Event struct {
	Type    string       `json:"event"`
	Date    string       `json:"date"`
	User    user         `json:"user"`
	Payload eventPayload `json:"payload"`
}

var hookHandler = func(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Ignoring", r.Method, "request.")
		return
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var t Event
	err := decoder.Decode(&t)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Received event \"", t.Type, "\".")

	switch t.Type {
	case "one-off-execution-created":
		triggerBDDBuild(t.Payload.Identifier)
	default:
		log.Println("No handler for ", t.Type, "implemented.")
	}

}

func main() {
	http.HandleFunc("/", hookHandler)

	log.Println("Webhook listener started ...")
	http.ListenAndServe(":3600", nil)
}
