package rest

type Resource struct {
	Method    string
	Path      string
	Unmarshal func([]byte) interface{}
	Handler   func(RequestContext) interface{}
}

type ResourceDefinition struct {
	Path      string
	Resources []Resource
}

type ResourceConfigurator interface {
	Configure(ResourceDefinition)
}
