package rest

type RequestHandler interface {
	NewRequestContext() RequestContext
}

type RequestContext struct {
	pathParamHandler  func(string) string
	queryParamHandler func(string) string
	entityHandler     func(interface{})
}

func (rc *RequestContext) QueryParam(name string) string {
	return rc.queryParamHandler(name)
}

func (rc *RequestContext) PathParam(name string) string {
	return rc.pathParamHandler(name)
}

func (rc *RequestContext) Entity(entity interface{}) {
	rc.entityHandler(entity)
}
