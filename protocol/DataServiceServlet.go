package protocol

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DataServiceServlet struct {
}

func (self *DataServiceServlet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r)
	self.handleSingleRequest(w, r)
}

func (self *DataServiceServlet) handleSingleRequest(w http.ResponseWriter, r *http.Request) {
	handler := self.getRestHandler(r)

	responseItem := handler.execute(r)
	// if errorcode >= 200 && < 400 then sendError
	enc := json.NewEncoder(w)
	enc.Encode(responseItem)
}

func (self *DataServiceServlet) getRestHandler(r *http.Request) RestHandler {
	return nil
}

func NewDataServiceServlet() *DataServiceServlet {
	return &DataServiceServlet{}
}
