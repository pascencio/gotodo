package rest

// Resource represent a single REST service operationg
type Resource interface {
	GetMethod() string
	GetPath() string
	Handler(RequestContext) interface{}
}

// ResourceDefinition represent a group of REST service resources
type ResourceDefinition struct {
	Path      string
	Resources []Resource
}

type ResourceConfigurator interface {
	Configure(ResourceDefinition)
}
