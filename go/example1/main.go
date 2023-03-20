// Basic example script for accessing a REMS REST API using go language.
// For more details on this script, see the README.md file.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Printf("Start: %v", time.Now())

	var apiUrl string = "https://rems-demo.rahtiapp.fi/api"
	var apiEnd string = "my-applications"
	var apiKey string = "55"
	var userId string = "RDapplicant1@funet.fi"

	// We need control over headers so we can't use the basic http.Get()
	// function which means we need to step up to the use of a client.
	client := &http.Client{}

	// Craft our message
	var url string = apiUrl + "/" + apiEnd
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-rems-api-key", apiKey)
	req.Header.Add("x-rems-user-id", userId)

	// Send our message
	log.Printf("Contacting REMS instance at %s", url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Houston, we have a problem: %s", err)
	}

	// What did we get back?
	log.Printf("Response Status: %s", resp.Status)
	log.Printf("Response Protocol: %s", resp.Proto)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Problem reading Response body: %s", err)
	}

	// body is a []byte so stringify for output
	fmt.Println(string(body))

	log.Printf("End: %v", time.Now())
}
