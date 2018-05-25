package todo

type TodoRepository interface {
	FindById(id interface{}) (*Todo, error)
	FindAll() (*[]Todo, error)
	Insert(todo *Todo) (*Todo, error)
	Update(todo *Todo) (*Todo, error)
	Delete(todo *Todo) (*Todo, error)
}
