package rest

// RequestHandler handle any http request
type RequestHandler interface {
	NewRequestContext() RequestContext
}

// RequestContext contains the context of http request
type RequestContext struct {
	pathParamHandler  func(string) string
	queryParamHandler func(string) string
	entityHandler     func(interface{})
}

// QueryParam return query params by name
func (rc *RequestContext) QueryParam(name string) string {
	return rc.queryParamHandler(name)
}

// PathParam return path params by name
func (rc *RequestContext) PathParam(name string) string {
	return rc.pathParamHandler(name)
}

// Entity return entity from http body
func (rc *RequestContext) Entity(entity interface{}) {
	rc.entityHandler(entity)
}
