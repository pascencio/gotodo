package todo

import "github.com/globalsign/mgo/bson"

// Todo domain for store todo tasks
type Todo struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" valid:"optional"`
	Title       string        `json:"title,omitempty" valid:"required,length(10|20)"`
	Description string        `json:"description,omitempty" valid:"required"`
}

// GetID get ID from todo domain
func (d Todo) GetID() interface{} {
	return d.ID
}

// SetID set ID to todo domain
func (d *Todo) SetID(ID interface{}) {
	d.ID = ID.(bson.ObjectId)
}
