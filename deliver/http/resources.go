package http

// Resource represent a single REST service operationg
type Resource interface {
	GetMethod() string
	GetPath() string
	Handler(RequestContext) (interface{}, error)
}

// ResourceDefinition represent a group of REST service resources
type ResourceDefinition struct {
	Path      string
	Resources []Resource
}

// ResourceConfigurator configure resource definitions
type ResourceConfigurator interface {
	Configure(ResourceDefinition)
}
