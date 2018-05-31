package mongo

import (
	"github.com/pascencio/gotodo/domain"
	"github.com/pascencio/gotodo/repository"
)

// TodoRepository for todo app
type TodoRepository struct {
	repository.Template
}

// FindByID find todo by id
func (r TodoRepository) FindByID(id interface{}) repository.Iterator {
	return r.Template.FindByID(id, "todo")
}

// FindAll find all todo
func (r TodoRepository) FindAll() repository.Iterator {
	return r.Template.FindAll("todo")
}

// Insert insert a todo
func (r TodoRepository) Insert(result domain.Domain) error {
	err := r.Template.Insert(result, "todo")
	return err
}

// Update update a todo
func (r TodoRepository) Update(result domain.Domain) error {
	err := r.Template.Update(result, "todo")
	return err
}

// Delete delete a todo
func (r TodoRepository) Delete(result domain.Domain) error {
	err := r.Template.Delete(result, "todo")
	return err
}
