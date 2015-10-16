package servlet

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/WitzHsiao/shindigo/social/service"
)

type RestResponseItem struct {
}

type RestHandler interface {
	handleItem(reqItem *service.RestRequestItem) *RestResponseItem
}

type DataServiceServlet struct {
	handlers map[string]RestHandler
}

func (self *DataServiceServlet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r)
	self.handleSingleRequest(w, r)
}

func (self *DataServiceServlet) handleSingleRequest(w http.ResponseWriter, r *http.Request) {
	reqItem := service.NewRestRequestItem(r)
	respItem, _ := self.handleRequestItem(reqItem)

	//responseItem := handler.execute(reqITem)
	// if errorcode >= 200 && < 400 then sendError
	enc := json.NewEncoder(w)
	enc.Encode(respItem)
}

func (self *DataServiceServlet) handleRequestItem(reqItem *service.RestRequestItem) (*RestResponseItem, error) {
	if handler, ok := self.handlers[reqItem.GetService()]; !ok {
		return nil, errors.New("The service " + reqItem.GetService() + " is not implemented")
	} else {
		return handler.handleItem(reqItem), nil
	}
}

//func (self *DataServiceServlet) getRestHandler(r *http.Request) RestHandler {
//return nil
//}

func NewDataServiceServlet() *DataServiceServlet {

	return &DataServiceServlet{}
}
