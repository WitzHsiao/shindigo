package protocol

type RestHandler interface {
	execute(r *RequestItem) ResponseItem
}
