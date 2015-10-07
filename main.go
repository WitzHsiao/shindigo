package main

import (
	"log"
	"net/http"

	"./social/opensocial/service"
)

func main() {
	http.Handle("/social/rest/", service.NewDataServiceServlet())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
