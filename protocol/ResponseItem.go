package protocol

type ResponseItem struct {
	ErrorCode    int
	ErrorMessage string
	Response     interface{}
}
