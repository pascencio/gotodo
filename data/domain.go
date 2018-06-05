package data

// Domain base interface for all domains
type Domain interface {
	GetID() interface{}
	SetID(interface{})
}
