package repository

import (
	"github.com/pascencio/gotodo/domain"
)

// Repository interface for data operations
type Repository interface {
	FindAll() Iterator
	FindByID(interface{}) Iterator
	Insert(domain.Domain) error
	Update(domain.Domain) error
	Delete(domain.Domain) error
}

// Iterator result collection iterator
type Iterator interface {
	Next(domain.Domain) bool
}

// Template interface for template of data operations
type Template interface {
	FindAll(string) Iterator
	FindByID(interface{}, string) Iterator
	Insert(domain.Domain, string) error
	Update(domain.Domain, string) error
	Delete(domain.Domain, string) error
}

// Connection database connection
type Connection interface {
	Close()
}

// ConnectionPool pool of database connection
type ConnectionPool interface {
	GetConnection() interface{}
}
