package rest

import (
	"fmt"
	"net/http"

	"github.com/pascencio/gotodo/domain"

	"github.com/pascencio/gotodo/repository"
)

const (
	methodInsert   string = "Insert"
	methodUpdate   string = "Update"
	methodDelete   string = "Delete"
	methodFindAll  string = "FindAll"
	methodFindByID string = "FindByID"
)

type deserializeEntity func(RequestContext) domain.Domain

type fetchDomain func(repository.Iterator) domain.Domain

type parseID func(string) interface{}

// CrudResource resource for CRUD operations
type CrudResource struct {
	Repository        repository.Repository
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
func (r CrudResource) Handler(c RequestContext) interface{} {
	switch r.CrudMethod {
	case methodInsert:
		result := r.DeserializeEntity(c)
		r.Repository.Insert(result)
		return result
	case methodUpdate:
		result := r.DeserializeEntity(c)
		r.Repository.Update(result)
		return result
	case methodDelete:
		ID := c.PathParam("id")
		result := r.DeserializeEntity(c)
		result.SetID(r.ParseID(ID))
		r.Repository.Delete(result)
		return result
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
		return results
	case methodFindByID:
		ID := c.PathParam("id")
		iterator := r.Repository.FindByID(ID)
		element := r.FetchDomain(iterator)
		return element
	default:
		panic(fmt.Errorf("Method not exists: [metho='%s']", r.CrudMethod))
	}
}

// NewCrudResourceDefinition create a CRUD resource
func NewCrudResourceDefinition(path string, r repository.Repository, d deserializeEntity, f fetchDomain, p parseID) ResourceDefinition {

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
