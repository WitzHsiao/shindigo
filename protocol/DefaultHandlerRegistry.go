package protocol

type DefaultHandlerRegistry struct {
}

func (self *DefaultHandlerRegistry) addHandlers(handlers []RestHandler) {
}

func (self *DefaultHandlerRegistry) getRestHandler(path string, method string) RestHandler {
	return nil
}
