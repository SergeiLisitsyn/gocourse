package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//creation maps with interface for formating JSON output
type Respon map[string]interface{}
type Head map[string]interface{}

func handler(w http.ResponseWriter, r *http.Request) {

	//create variable for Respose
	resp := Respon{
		"host":        r.Host,
		"user_agent":  r.UserAgent(),
		"request_uri": r.RequestURI,
		"headers": Head{
			"Accept":     r.Header["Accept"],
			"User-Agent": r.Header["User-Agent"],
		},
	}
	// create JSON
	responseJSON, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}

	// Response message
	fmt.Fprint(w, string(responseJSON))
}

func main() {

	const port = 8080

	fmt.Printf("Launching server on port: %d \n\n", port)
	// set handler for rout '/'
	http.HandleFunc("/", handler)
	// start server without end
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))

}
