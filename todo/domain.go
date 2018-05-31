package todo

import "github.com/globalsign/mgo/bson"

// Todo domain for store todo tasks
type Todo struct {
	ID          bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
}

// GetID get ID from todo domain
func (d Todo) GetID() interface{} {
	return d.ID
}

// SetID set ID to todo domain
func (d *Todo) SetID(ID interface{}) {
	d.ID = ID.(bson.ObjectId)
}
