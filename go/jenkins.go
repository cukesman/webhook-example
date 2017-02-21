package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type jenkinsParameters struct {
	Parameters []parameter `json:"parameter"`
}

type parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func triggerBDDBuild(id string) {
	parameters := []parameter{parameter{Name: "CUKESMAN_ONEOFF_FEATURE_ID", Value: id}}
	jenkinsParameters := &jenkinsParameters{Parameters: parameters}

	jenkinsParameterJSON, _ := json.Marshal(jenkinsParameters)

	jenkinsURL := os.Getenv("JENKINS_URL")
	if jenkinsURL == "" {
		jenkinsURL = "http://localhost:38080"
	}

	jenkinsBuild := os.Getenv("JENKINS_BUILD")
	if jenkinsBuild == "" {
		jenkinsBuild = "param-test"
	}

	jenkinsToken := os.Getenv("JENKINS_TOKEN")
	if jenkinsToken == "" {
		jenkinsToken = "ffe0d2c93e8293867bad65ff364bf8c2"
	}

	buildTriggerURL := fmt.Sprintf("%s/job/%s/build", jenkinsURL, jenkinsBuild)
	log.Println("Trigger Jenkins Build for One Off Execution", id, "at", buildTriggerURL)

	formData := url.Values{}
	formData.Set("token", jenkinsToken)
	formData.Add("json", string(jenkinsParameterJSON))

	response, err := http.PostForm(buildTriggerURL, formData)

	if response.StatusCode != 201 {
		panic(fmt.Sprintf("Received status code %s from Jenkins (expected 201).", response.Status))
	}

	if err != nil {
		panic(err)
	}
}
