package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/dtoebe/collab-md/internal/methodcheck"
	"github.com/dtoebe/collab-md/internal/middleware"
)

func main() {

	var host, port string
	flag.StringVar(&host, "h", "", "host to listen on")
	flag.StringVar(&port, "p", "3080", "port to listen on")

	flag.Parse()

	addr := net.JoinHostPort(host, port)
	server := http.Server{
		Addr: addr,
	}

	hub := NewHub("http", addr)

	http.Handle("/documents/create/",
		middleware.Assign(
			http.HandlerFunc(docCreateHandler),
			methodcheck.AllowedMethods("POST"),
			hub.NewInstanceMiddleware(),
		),
	)
	http.Handle("/documents/edit/",
		middleware.Assign(
			http.HandlerFunc(docEditHandler),
			methodcheck.AllowedMethods("GET"),
		),
	)
	http.HandleFunc("/", indexHandler)

	log.Println("listening on", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
