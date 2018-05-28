package todo

// Todo domain for store todo tasks
type Todo struct {
	ID          interface{} `json:"-"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
}

// GetID get ID from todo domain
func (d Todo) GetID() interface{} {
	return d.ID
}
