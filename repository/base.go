package repository

type Repository struct {
	Template RepositoryTemplate
}

type RepositoryTemplate interface {
	SetConnection(connection Connection)
	FindAll(interface{}, string) error
	FindByID(interface{}, interface{}, string) error
	Insert(interface{}, string) error
	Update(interface{}, string) error
	Delete(interface{}, string) error
}

type Connection interface {
	Close() error
}

type ConnectionPool interface {
	GetConnection() Connection
	Start() error
}
