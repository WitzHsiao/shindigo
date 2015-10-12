package main

import (
	"log"
	"net/http"

	"./protocol"
)

func main() {
	http.Handle("/social/rest/", protocol.NewDataServiceServlet())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
