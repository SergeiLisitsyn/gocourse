package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var name string
var address string

type Respon map[string]interface{}

func handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

		// GET return HTML page
		http.ServeFile(w, r, "form.html")

	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//Creation of a cookie
		name = r.FormValue("name")
		address = r.FormValue("address")

		expire := time.Now().Add(240 * time.Hour)
		cookie := &http.Cookie{
			Name:    "token",
			Value:   fmt.Sprintf("%s:%s", name, address),
			MaxAge:  300,
			Expires: expire,
		}

		http.SetCookie(w, cookie)

		w.WriteHeader(200)

		//create variable for Respose
		resp := Respon{
			"token":      cookie.Value,
			"createdAtt": time.Now().Format("2006-01-02 3:4:5 pm"),
			"expiredAt":  cookie.Expires.Format("2006-01-02 3:4:5 pm"),
		}
		// create JSON
		json_data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			log.Fatalln(err)
		}
		// token is sent with restAPI
		res, err := http.Post("http://127.0.0.1:8082/tokens", "application/json",
			bytes.NewBuffer(json_data))
		if err != nil {
			log.Fatal(err)
		}

		var re map[string]interface{}

		json.NewDecoder(res.Body).Decode(&re)
		//Print cookie

		fmt.Println(cookie)
		http.ServeFile(w, r, "form.html")
	default:
		fmt.Fprintf(w, "404 not found.")
	}
}

func main() {
	var port int = 8080
	fmt.Printf("Server started at port %d\n", port)
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
