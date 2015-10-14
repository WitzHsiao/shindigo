package protocol

import "net/http"

type RestHandler interface {
	//execute(r *RequestItem) ResponseItem
	execute(r *http.Request) *ResponseItem
}
