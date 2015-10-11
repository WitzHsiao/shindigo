package protocol

type RequestItem interface {
	getAppId() string
	getParameter(paramName string, defaultValue string) string
	getListParameter(paramName string) []string
}
