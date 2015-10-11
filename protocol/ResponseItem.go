package protocol

type ResponseItem struct {
	ErrorCode    int
	ErrorMessage string
	Response     interface{}
}

//func NewResponseItem(errorCode int, errorMessage string, response interface{}) *ResponseItem {
//return &ResponseItem{
//ErrorCode:    errorCode,
//ErrorMessage: errorMessage,
//Response:     response,
//}
//}
