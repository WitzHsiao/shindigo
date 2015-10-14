package protocol

import (
	"net/http"
	"testing"
)

var registry *DefaultHandlerRegistry

func setup() {
	registry = &DefaultHandlerRegistry{}
	registry.addHandlers([]RestHandler{&TestHandler{}})
}

func TestGetHandlerRest(t *testing.T) {
	setup()
	handler := registry.getRestHandler("/test", "GET")
	if handler == nil {
		t.Fail()
	}
}

type TestHandler struct {
}

func (testHandler *TestHandler) execute(r *http.Request) *ResponseItem {
	return nil
}
