package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type event struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type store struct {
	events map[string][]event
	m      sync.RWMutex
}

// Global vars
var (
	s = store{
		events: map[string][]event{},
		m:      sync.RWMutex{},
	}
	port = flag.Int("p", 8080, "Port httpd should listen on")
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "JSON API Mock for testing\n")
}

func ingest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	event := &event{}
	fmt.Printf("Got new event from %s\n", ps.ByName("uuid"))
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading the JSON body\n")
		return
	}

	err = json.Unmarshal(jsn, event)
	if err != nil {
		fmt.Fprintf(w, "Error unmarshaling JSON payload\n")
		return
	}

	s.m.Lock()
	s.events[ps.ByName("uuid")] = append(s.events[ps.ByName("uuid")], *event)
	s.m.Unlock()
	fmt.Fprintf(w, "Success\n")

}

func getevents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	s.m.RLock()
	json, err := json.Marshal(s.events[ps.ByName("uuid")])

	if err != nil {
		fmt.Fprintf(w, "Error marshaling JSON payload\n")
		return
	}

	fmt.Fprintf(w, "%s\n", json)
	s.m.RUnlock()
}

func main() {
	// Uncomment to generate UUID, e.g. for testing
	/*uuid, _ := uuid.NewRandom()
	fmt.Printf("Generated UUID: %v\n", uuid)
	*/

	// Command line parameters
	flag.Parse()

	// httprouter initialization
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/events/:uuid", getevents)
	router.POST("/events/ingest/:uuid", ingest)

	// Run or panic
	fmt.Printf("Starting http daemon on port %d...\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
