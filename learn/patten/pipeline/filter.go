package pipeline

type Request interface{}

type Response interface{}

type Filter interface {
	Process(req Request) (Response, error)
}
