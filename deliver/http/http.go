package http

// RequestHandler handle any http request
type RequestHandler interface {
	NewRequestContext() RequestContext
}

type entityHandler func(interface{}) error

type queryParamHandler func(string) string

type pathParamHandler func(string) string

// RequestContext contains the context of http request
type RequestContext struct {
	pathParamHandler  pathParamHandler
	queryParamHandler queryParamHandler
	entityHandler     entityHandler
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
func (rc *RequestContext) Entity(entity interface{}) error {
	return rc.entityHandler(entity)
}

// ValidationError validation error
type ValidationError interface {
	Details() map[string]string
}
