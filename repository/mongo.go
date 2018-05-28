package repository

import (
	"time"

	"github.com/pascencio/gotodo/domain"
	"github.com/spf13/viper"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

// MongoRepositoryTemplate for mongodb operations
type MongoRepositoryTemplate struct {
	connection MongoConnection
}

// FindAll get all data from specific collection
func (r *MongoRepositoryTemplate) FindAll(result []domain.Domain, name string) error {
	defer r.connection.Close()
	err := r.getCollection(name).Find(nil).All(&result)
	return err
}

// FindByID get single document from specific collection
func (r MongoRepositoryTemplate) FindByID(id interface{}, result domain.Domain, name string) error {
	defer r.connection.Close()
	err := r.getCollection(name).FindId(id).One(result)
	return err
}

// Insert single document on specific collection
func (r MongoRepositoryTemplate) Insert(result domain.Domain, name string) error {
	defer r.connection.Close()
	err := r.getCollection(name).Insert(&result)
	return err
}

// Update single document on specific collection
func (r *MongoRepositoryTemplate) Update(result domain.Domain, name string) error {
	defer r.connection.Close()
	err := r.getCollection(name).UpdateId(result.GetID(), result)
	return err
}

// Delete single document on specific collection
func (r *MongoRepositoryTemplate) Delete(result domain.Domain, name string) error {
	defer r.connection.Close()
	err := r.getCollection(name).RemoveId(result.GetID())
	return err
}

func (r *MongoRepositoryTemplate) getCollection(name string) *mgo.Collection {
	session := r.connection.sessionCopy
	return session.DB(r.connection.database).C(name)
}

// SetConnection assign single connection to template
func (r *MongoRepositoryTemplate) SetConnection(connection MongoConnection) {
	r.connection = connection
}

// MongoConnection mongodb connection
type MongoConnection struct {
	sessionCopy *mgo.Session
	database    string
}

// Close a single sessionCopy
func (c MongoConnection) Close() error {
	c.sessionCopy.Close()
	return nil
}

// MongoConnectionPool connection pool for mongodb
type MongoConnectionPool struct {
	mongoSession *mgo.Session
	database     string
}

// GetConnection return single connection from pool
func (p MongoConnectionPool) GetConnection() Connection {
	return MongoConnection{sessionCopy: p.mongoSession.Copy(), database: p.database}
}

// Start inite the connection pool
func (p MongoConnectionPool) Start() error {
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
