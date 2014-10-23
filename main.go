package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"strconv"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 9500, "The port to bind the server to.")

	http.HandleFunc("/", handleRoot)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	var requestPath = strings.ToLower(strings.Trim(r.URL.Path, " /"))

	var redirectUrl string

	if requestPath == "" {
		redirectUrl = "http://projects.oscelot.org/gf/project/jshack/"
	} else {
		redirectUrl = "http://projects.oscelot.org/gf/project/jshack/wiki/?pagename=" + requestPath
	}
	
	log.Printf("Redirecting from: '%s' to: '%s'", r.URL.String(), redirectUrl)
	http.Redirect(w, r, redirectUrl, 302)
}