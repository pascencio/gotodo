package rest

import (
	"fmt"
	"net/http"

	"github.com/pascencio/gotodo/domain"

	"github.com/pascencio/gotodo/repository"
)

// CrudRequestHandler handle integration between domain variable declaration, REST request and Repository
type CrudRequestHandler struct {
	Repository     repository.Repository
	Method         string
	RequestContext RequestContext
	Domain         domain.Domain
}

// Handle asociate HTTP method with CRUD repository operation
func (c *CrudRequestHandler) Handle(d domain.Domain) {
	switch c.Method {
	case http.MethodPost:
		c.RequestContext.Entity(d)
		//c.Repository.Insert(&d)
		c.Domain = d
	default:
		panic(fmt.Errorf("Method not exists: [metho='%s']", c.Method))
	}
}

// CrudResource resource for CRUD operations
type CrudResource struct {
	Repository     repository.Repository
	path           string
	method         string
	AllocateDomain func(*CrudRequestHandler)
}

// GetPath path of resource
func (r CrudResource) GetPath() string {
	return r.path
}

// GetMethod method of resource
func (r CrudResource) GetMethod() string {
	return r.method
}

// Handler handle a request to the resource
func (r CrudResource) Handler(context RequestContext) interface{} {
	crudHandler := CrudRequestHandler{
		Repository:     r.Repository,
		RequestContext: context,
		Method:         r.GetMethod(),
	}

	r.AllocateDomain(&crudHandler)

	return crudHandler.Domain
}

// NewCrudResourceDefinition create a CRUD resource
func NewCrudResourceDefinition(p string, r repository.Repository, d func(*CrudRequestHandler)) ResourceDefinition {

	resources := []Resource{
		CrudResource{
			Repository:     r,
			AllocateDomain: d,
			method:         http.MethodGet,
		},
		CrudResource{
			Repository:     r,
			AllocateDomain: d,
			method:         http.MethodGet,
			path:           ":id",
		},
		CrudResource{
			Repository:     r,
			AllocateDomain: d,
			method:         http.MethodPost,
		},
		CrudResource{
			Repository:     r,
			AllocateDomain: d,
			method:         http.MethodPut,
			path:           ":id",
		},
		CrudResource{
			Repository:     r,
			AllocateDomain: d,
			method:         http.MethodDelete,
			path:           ":id",
		},
	}

	definition := ResourceDefinition{
		Path:      p,
		Resources: resources,
	}

	return definition
}
