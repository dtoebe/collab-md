package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// methodcheck.NotImplemented(w)
	fmt.Fprintln(w, "Index")
	return
}

func docCreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Document Created")
}

func docEditHandler(w http.ResponseWriter, r *http.Request) {
	urlSplit := strings.Split(r.URL.Path, "/")
	log.Println("Edit requested")
	fmt.Fprintln(w, "Connection-Type:",
		r.Header.Get("Connection-Type"),
		"ID:", urlSplit[len(urlSplit)-1],
	)
}

func docdeleteHandler(w http.ResponseWriter, r *http.Request) {

}
