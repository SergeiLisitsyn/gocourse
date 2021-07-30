package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var name string
var address string

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

		name = r.FormValue("name")

		address = r.FormValue("address")

		fmt.Fprintf(w, "%s from %s\n", name, address)
		fmt.Printf("%s from %s\n", name, address)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func docHandler(w http.ResponseWriter, r *http.Request) {

	expire := time.Now().Add(200 * time.Minute)
	cookie := &http.Cookie{
		Name:    "token#1",
		Value:   fmt.Sprintf("%s:%s", name, address),
		MaxAge:  300,
		Expires: expire,
	}

	http.SetCookie(w, cookie)

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("%s Get Successful\n", cookie.Name)))

	//Print Cookie

	fmt.Println(cookie)

	return
}

func main() {
	var port int = 8080
	fmt.Printf("Server started at port %d\n", port)
	http.HandleFunc("/", handler)
	docHandler := http.HandlerFunc(docHandler)
	http.Handle("/doc", docHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
