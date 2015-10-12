package protocol

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const HOST = "http://0.0.0.0"
const X_HTTP_METHOD_OVERRIDE = "X-HTTP-Method-Override"

func TestUriRecognition(t *testing.T) {
	pathInfo := "/test/5/@self"
	verifyHandlerWasFoundForPathInfo(t, pathInfo, "POST", "POST")
}

func verifyHandlerWasFoundForPathInfo(t *testing.T, pathInfo, actualMethod, overrideMethod string) {
	req, err := getFakeRequest(pathInfo, actualMethod, overrideMethod)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	w.Header().Add("Content-Type", "application/json")

	servlet := NewDataServiceServlet()
	servlet.ServeHTTP(w, req)
	_, err = w.Write([]byte("GET_RESPONSE"))
	if err != nil {
		t.Error(err)
	}
}

func getFakeRequest(pathInfo string, actualMethod string, overrideMethod string) (*http.Request, error) {
	fakeReq, err := http.NewRequest(actualMethod, HOST+pathInfo, nil)
	if err != nil {
		return nil, err
	}
	fakeReq.Header.Add(X_HTTP_METHOD_OVERRIDE, overrideMethod)
	return fakeReq, nil
}
