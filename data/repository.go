package data

// Repository interface for data operations
type Repository interface {
	FindAll() Iterator
	FindByID(interface{}) Iterator
	Insert(Domain) error
	Update(Domain) error
	Delete(Domain) error
}

// Iterator result collection iterator
type Iterator interface {
	Next(Domain) bool
}

// Template interface for template of data operations
type Template interface {
	FindAll(string) Iterator
	FindByID(interface{}, string) Iterator
	Insert(Domain, string) error
	Update(Domain, string) error
	Delete(Domain, string) error
}

// Connection database connection
type Connection interface {
	Close()
}

// ConnectionPool pool of database connection
type ConnectionPool interface {
	GetConnection() interface{}
}
