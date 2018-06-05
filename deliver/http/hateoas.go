package http

import (
	"fmt"
	"net/http"

	"github.com/pascencio/gotodo/data"
)

const (
	methodInsert   string = "Insert"
	methodUpdate   string = "Update"
	methodDelete   string = "Delete"
	methodFindAll  string = "FindAll"
	methodFindByID string = "FindByID"
)

type deserializeEntity func(RequestContext) (data.Domain, error)

type fetchDomain func(data.Iterator) data.Domain

type parseID func(string) interface{}

// CrudResource resource for CRUD operations
type CrudResource struct {
	Repository        data.Repository
	CrudMethod        string
	path              string
	method            string
	DeserializeEntity deserializeEntity
	FetchDomain       fetchDomain
	ParseID           parseID
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
func (r CrudResource) Handler(c RequestContext) (interface{}, error) {
	switch r.CrudMethod {
	case methodInsert:
		result, err := r.DeserializeEntity(c)
		if err != nil {
			return nil, err
		}
		r.Repository.Insert(result)
		return result, nil
	case methodUpdate:
		result, err := r.DeserializeEntity(c)
		if err != nil {
			return nil, err
		}
		r.Repository.Update(result)
		return result, nil
	case methodDelete:
		ID := c.PathParam("id")
		iterator := r.Repository.FindByID(ID)
		element := r.FetchDomain(iterator)
		r.Repository.Delete(element)
		return element, nil
	case methodFindAll:
		iterator := r.Repository.FindAll()
		var results []interface{}
		for {
			element := r.FetchDomain(iterator)
			if element == nil {
				break
			}
			results = append(results, element)
		}
		return results, nil
	case methodFindByID:
		ID := c.PathParam("id")
		iterator := r.Repository.FindByID(ID)
		element := r.FetchDomain(iterator)
		return element, nil
	default:
		return nil, fmt.Errorf("Method not exists: [metho='%s']", r.CrudMethod)
	}
}

// NewCrudResourceDefinition create a CRUD resource
func NewCrudResourceDefinition(path string, r data.Repository, d deserializeEntity, f fetchDomain, p parseID) ResourceDefinition {

	resources := []Resource{
		CrudResource{
			Repository:        r,
			DeserializeEntity: d,
			FetchDomain:       f,
			method:            http.MethodGet,
			CrudMethod:        methodFindAll,
		},
		CrudResource{
			Repository:        r,
			DeserializeEntity: d,
			FetchDomain:       f,
			method:            http.MethodGet,
			path:              ":id",
			CrudMethod:        methodFindByID,
			ParseID:           p,
		},
		CrudResource{
			Repository:        r,
			DeserializeEntity: d,
			FetchDomain:       f,
			method:            http.MethodPost,
			CrudMethod:        methodInsert,
		},
		CrudResource{
			Repository:        r,
			DeserializeEntity: d,
			FetchDomain:       f,
			method:            http.MethodPut,
			CrudMethod:        methodUpdate,
		},
		CrudResource{
			Repository:        r,
			DeserializeEntity: d,
			FetchDomain:       f,
			method:            http.MethodDelete,
			path:              ":id",
			CrudMethod:        methodDelete,
			ParseID:           p,
		},
	}

	definition := ResourceDefinition{
		Path:      path,
		Resources: resources,
	}

	return definition
}
