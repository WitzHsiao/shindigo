package service

import (
	"net/http"
	"net/url"
	"strings"
)

type RestRequestItem struct {
	request *http.Request
}

func NewRestRequestItem(r *http.Request) *RestRequestItem {
	return &RestRequestItem{
		request: r,
	}
}

func (self *RestRequestItem) GetParameters() url.Values {
	return self.request.URL.Query()
}

func (self *RestRequestItem) GetService() string {
	reqUri := self.request.URL.RequestURI()
	reqUri = strings.TrimPrefix(reqUri, "/")
	strs := strings.Split(reqUri, "/")
	return strs[0]
}
