package protocol

import (
	"fmt"
	"net/http"
)

type DataServiceServlet struct {
}

func (dataServiceServlet *DataServiceServlet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r)
}

func NewDataServiceServlet() *DataServiceServlet {
	return &DataServiceServlet{}
}
