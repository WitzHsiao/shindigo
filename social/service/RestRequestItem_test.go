package service

import (
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var urlString string

func setup() {
	params := map[string]string{
		"oauth_nonce":            "10075052d8a3cd0087d11346edba8f1f",
		"oauth_timestamp":        "1242011332",
		"oauth_consumer_key":     "consumerKey",
		"fields":                 "gender,name,phone",
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_signature":        "wDcyXTBqhxW70G+ddZtw7zPVGyE=",
	}

	var urlencodedParams []string
	for key, value := range params {
		urlencodedParams = append(urlencodedParams, key+"="+url.QueryEscape(value))
	}
	urlString = "/people/1/@self?" + strings.Join(urlencodedParams, "&")
}

func TestCreateWithRequest(t *testing.T) {
	setup()

	urlURL, err := url.Parse(urlString)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		t.Error(err)
	}
	reqItem := NewRestRequestItem(req)
	actual := reqItem.GetParameters()
	expect := urlURL.Query()
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("\n%#v \nis not equals to \n%#v", expect, actual)
	}

}

func TestGetService(t *testing.T) {
	setup()

	req, err := http.NewRequest("GET", "http://localhost:9999"+urlString, nil)
	if err != nil {
		t.Error(err)
	}
	reqItem := NewRestRequestItem(req)
	if "people" != reqItem.GetService() {
		t.Errorf("%s is not equal to %s", reqItem.GetService(), "people")
	}
}
