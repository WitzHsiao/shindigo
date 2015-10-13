package main

import (
	"log"
	"net/http"

	"github.com/WitzHsiao/shindigo/protocol"
)

func main() {
	http.Handle("/social/rest/", protocol.NewDataServiceServlet())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
