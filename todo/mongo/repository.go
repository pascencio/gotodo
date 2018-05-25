package mongo

import (
	"github.com/pascencio/gotodo/repository"
	"github.com/pascencio/gotodo/todo"
)

type TodoRepository struct {
	repository.Repository
}

func (r TodoRepository) FindById(id interface{}) (*todo.Todo, error) {
	result := todo.Todo{}
	err := r.Template.FindById(id, &result, "todo")
	return &result, err
}
func (r TodoRepository) FindAll() (*[]todo.Todo, error) {
	domains := []todo.Todo{}
	err := r.Template.FindAll(&domains, "todo")
	return &domains, err
}
func (r TodoRepository) Insert(result *todo.Todo) (*todo.Todo, error) {
	err := r.Template.Insert(&result, "todo")
	return result, err
}
func (r TodoRepository) Update(result *todo.Todo) (*todo.Todo, error) {
	err := r.Template.Update(&result, "todo")
	return result, err
}
func (r TodoRepository) Delete(result *todo.Todo) (*todo.Todo, error) {
	err := r.Template.Delete(&result, "todo")
	return result, err
}
