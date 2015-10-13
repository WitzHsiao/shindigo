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

func TestCount(t *testing.T) {
	req.SetParameter("count", nil)
	actual, _ := req.GetCount()
	if !(DEFAULT_COUNT == actual) {
		t.Fail()
	}

	req.SetParameter("count", "5")
	actual, _ = req.GetCount()
	if !(5 == actual) {
		t.Fail()
	}
}

func TestStartIndex(t *testing.T) {
	req.SetParameter("startIndex", nil)
	actual, _ := req.GetStartIndex()
	if !(DEFAULT_START_INDEX == actual) {
		t.Fail()
	}

	req.SetParameter("startIndex", "5")
	actual, _ = req.GetStartIndex()
	if !(5 == actual) {
		t.Fail()
	}
}
