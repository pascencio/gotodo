package todo

import (
	"github.com/pascencio/gotodo/data"
)

// Repository abstract todo repository
type Repository interface {
	FindById(id interface{}) data.Iterator
	FindAll() data.Iterator
	Insert(data.Domain) error
	Update(data.Domain) error
	Delete(data.Domain) error
}
