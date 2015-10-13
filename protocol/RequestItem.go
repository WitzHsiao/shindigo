package protocol

import (
	"fmt"
	"strconv"
	"strings"
)

// Common OpenSocial API fields
const APP_ID = "appId"
const START_INDEX = "startIndex"
const COUNT = "count"
const SORT_BY = "sortBy"
const SORT_ORDER = "sortOrder"
const FILTER_BY = "filterBy"
const FILTER_OPERATION = "filterOp"
const FILTER_VALUE = "filterValue"
const FIELDS = "fields" // Opensocial defaults
const DEFAULT_START_INDEX = 0
const DEFAULT_COUNT = 20
const APP_SUBSTITUTION_TOKEN = "@app"

type RequestItem struct {
	parameters map[string]interface{}
}

func NewRequestItem() *RequestItem {
	return &RequestItem{
		parameters: make(map[string]interface{}),
	}
}

func (self *RequestItem) GetListParameter(paramName string) []string {
	param, ok := self.parameters[paramName]
	if !ok {
		return []string{}
	}
	if p, ok := param.(string); ok && strings.Index(p, ",") != -1 {
		listParam := strings.Split(p, ",")
		self.parameters[paramName] = listParam
		return listParam
	} else if p, ok := param.([]string); ok {
		return p
		// TODO: JSONArray? List<?> ?
	} else {
		return []string{fmt.Sprint(param)}
	}
}

func (self *RequestItem) SetParameter(paramName string, paramValue interface{}) {
	if paramValue == nil {
		return
	}
	if p, ok := paramValue.([]string); ok {
		if len(p) == 1 {
			self.parameters[paramName] = p[0]
		} else {
			self.parameters[paramName] = p
		}
	} else if p, ok := paramValue.(string); ok {
		if len(p) > 0 {
			self.parameters[paramName] = p
		}
	} else {
		self.parameters[paramName] = paramValue
	}
}

func (self *RequestItem) GetAppId() string {
	appId := self.GetParameter(APP_ID)
	if appId != "" && appId == APP_SUBSTITUTION_TOKEN {
		// TODO: return token.getAppId()
		return "TODO"
	} else {
		return appId
	}
}

func (self *RequestItem) GetParameter(paramName string) string {
	param, ok := self.parameters[paramName]
	if !ok {
		return ""
	}
	if p, ok := param.([]interface{}); ok {
		if len(p) == 0 {
			return ""
		} else {
			param = p[0]
		}
	}
	return fmt.Sprint(param)
}

func (self *RequestItem) GetCount() (int, error) {
	count := self.GetParameter(COUNT)
	if count == "" {
		return DEFAULT_COUNT, nil
	} else {
		i, err := strconv.Atoi(count)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
}
