package mongo

import (
	"github.com/pascencio/gotodo/repository"
	"github.com/pascencio/gotodo/todo"
)

type MongoTodoRepository struct {
	repository.Repository
}

func (r MongoTodoRepository) FindByID(id interface{}) (*todo.Todo, error) {
	result := todo.Todo{}
	err := r.Template.FindByID(id, &result, "todo")
	return &result, err
}
func (r MongoTodoRepository) FindAll() (*[]todo.Todo, error) {
	domains := []todo.Todo{}
	err := r.Template.FindAll(&domains, "todo")
	return &domains, err
}
func (r MongoTodoRepository) Insert(result *todo.Todo) (*todo.Todo, error) {
	err := r.Template.Insert(&result, "todo")
	return result, err
}
func (r MongoTodoRepository) Update(result *todo.Todo) (*todo.Todo, error) {
	err := r.Template.Update(&result, "todo")
	return result, err
}
func (r MongoTodoRepository) Delete(result *todo.Todo) (*todo.Todo, error) {
	err := r.Template.Delete(&result, "todo")
	return result, err
}
