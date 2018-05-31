package repository

import (
	"time"

	"github.com/pascencio/gotodo/domain"
	"github.com/spf13/viper"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

// CollectionIterator iterates over mongodb collection
type CollectionIterator struct {
	Iter *mgo.Iter
}

// Next fetch single element. Return true if has more data to fetch.
func (i CollectionIterator) Next(result domain.Domain) bool {
	return i.Iter.Next(result)
}

// MongoRepositoryTemplate for mongodb operations
type MongoRepositoryTemplate struct {
	ConnectionPool ConnectionPool
}

// FindAll get all data from specific collection
func (r MongoRepositoryTemplate) FindAll(name string) Iterator {
	connection := r.ConnectionPool.GetConnection().(MongoConnection)
	defer connection.Close()
	iterator := CollectionIterator{
		Iter: r.getCollection(connection, name).Find(nil).Iter(),
	}
	return iterator
}

// FindByID get single document from specific collection
func (r MongoRepositoryTemplate) FindByID(id interface{}, name string) Iterator {
	connection := r.ConnectionPool.GetConnection().(MongoConnection)
	defer connection.Close()
	_id := r.buildID(id)
	iterator := CollectionIterator{
		Iter: r.getCollection(connection, name).FindId(_id).Iter(),
	}

	return iterator
}

// Insert single document on specific collection
func (r MongoRepositoryTemplate) Insert(result domain.Domain, name string) error {
	connection := r.ConnectionPool.GetConnection().(MongoConnection)
	defer connection.Close()
	err := r.getCollection(connection, name).Insert(&result)
	return err
}

// Update single document on specific collection
func (r MongoRepositoryTemplate) Update(result domain.Domain, name string) error {
	connection := r.ConnectionPool.GetConnection().(MongoConnection)
	defer connection.Close()
	err := r.getCollection(connection, name).UpdateId(result.GetID(), result)
	return err
}

// Delete single document on specific collection
func (r MongoRepositoryTemplate) Delete(result domain.Domain, name string) error {
	connection := r.ConnectionPool.GetConnection().(MongoConnection)
	defer connection.Close()
	err := r.getCollection(connection, name).RemoveId(result.GetID())
	return err
}

func (r MongoRepositoryTemplate) getCollection(connection MongoConnection, name string) *mgo.Collection {
	return connection.sessionCopy.DB(connection.database).C(name)
}

func (r MongoRepositoryTemplate) buildID(id interface{}) bson.ObjectId {
	return bson.ObjectIdHex(id.(string))
}

// SetConnection assign single connection to template
func (r *MongoRepositoryTemplate) SetConnection(connection ConnectionPool) {
	r.ConnectionPool = connection
}

//MongoConnection session instance
type MongoConnection struct {
	sessionCopy *mgo.Session
	database    string
}

//Close mongodb session copy
func (c MongoConnection) Close() {
	c.sessionCopy.Close()
}

// MongoConnectionPool connection pool for mongodb
type MongoConnectionPool struct {
	mongoSession *mgo.Session
	database     string
}

// GetConnection a copy of mongo sesion
func (p MongoConnectionPool) GetConnection() interface{} {
	mongoConnection := MongoConnection{
		database:    p.database,
		sessionCopy: p.mongoSession.Copy(),
	}

	return mongoConnection
}

// Start inite the connection pool
func (p *MongoConnectionPool) Start() error {
	addresses := viper.GetStringSlice("mongodb.addresses")
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    addresses,
		Timeout:  time.Duration(viper.GetInt("mongodb.timeout")) * time.Second,
		Database: viper.GetString("mongodb.database"),
		Username: viper.GetString("mongodb.username"),
		Password: viper.GetString("mongodb.password"),
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		log.WithFields(log.Fields{
			"addresses": addresses,
			"database":  viper.GetString("mongodb.database"),
		}).Fatal("Error al iniciar sesión en la base de datos")
		return err
	}

	p.mongoSession = mongoSession
	p.database = viper.GetString("mongodb.database")

	log.WithFields(log.Fields{
		"addresses": addresses,
		"database":  viper.GetString("mongodb.database"),
	}).Info("Conexión realizada correctamente")

	return nil
}
