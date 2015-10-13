package protocol

import (
	"reflect"
	"testing"
)

var req *RequestItem = NewRequestItem()

func TestParseCommaSeparetedList(t *testing.T) {
	req.SetParameter("fields", "huey,dewey,louie")
	expect := []string{"huey", "dewey", "louie"}
	actual := req.GetListParameter("fields")
	if !reflect.DeepEqual(expect, actual) {
		t.Fail()
	}
}

func TestGetAppId(t *testing.T) {
	req.SetParameter("appId", "100")
	if !("100" == req.GetAppId()) {
		t.Fail()
	}

	// TODO: test with fake gadget token
}

func TestStartIndex(t *testing.T) {
	req.SetParameter("count", nil)
	if !(DEFAULT_COUNT == req.GetCount()) {
		t.Fail()
	}

	req.SetParameter("count", "5")
	if !(5 == req.GetCount()) {
		t.Fail()
	}
}
