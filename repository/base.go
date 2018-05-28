package repository

import (
	"github.com/pascencio/gotodo/domain"
)

// Repository interface for data operations
type Repository interface {
	FindAll(*[]domain.Domain) error
	FindByID(interface{}, *domain.Domain) error
	Insert(*domain.Domain) error
	Update(*domain.Domain) error
	Delete(*domain.Domain) error
}

// Template interface for template of data operations
type Template interface {
	SetConnection(connection Connection)
	FindAll(interface{}, string) error
	FindByID(interface{}, interface{}, string) error
	Insert(interface{}, string) error
	Update(interface{}, string) error
	Delete(interface{}, string) error
}

// Connection database connection
type Connection interface {
	Close() error
}

// ConnectionPool pool of database connection
type ConnectionPool interface {
	GetConnection() Connection
	Start() error
}
