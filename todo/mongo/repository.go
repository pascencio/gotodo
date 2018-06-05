package mongo

import "github.com/pascencio/gotodo/data"

// TodoRepository for todo app
type TodoRepository struct {
	data.Template
}

// FindByID find todo by id
func (r TodoRepository) FindByID(id interface{}) data.Iterator {
	return r.Template.FindByID(id, "todo")
}

// FindAll find all todo
func (r TodoRepository) FindAll() data.Iterator {
	return r.Template.FindAll("todo")
}

// Insert insert a todo
func (r TodoRepository) Insert(result data.Domain) error {
	err := r.Template.Insert(result, "todo")
	return err
}

// Update update a todo
func (r TodoRepository) Update(result data.Domain) error {
	err := r.Template.Update(result, "todo")
	return err
}

// Delete delete a todo
func (r TodoRepository) Delete(result data.Domain) error {
	err := r.Template.Delete(result, "todo")
	return err
}
