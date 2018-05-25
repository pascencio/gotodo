package repository

import (
	"time"

	"github.com/pascencio/gotodo/domain"
	"github.com/spf13/viper"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

type MongoRepositoryTemplate struct {
	connection MongoConnection
}

func (r *MongoRepositoryTemplate) FindAll(result interface{}, name string) error {
	defer r.connection.Close()
	err := r.getCollection(name).Find(nil).All(&result)
	return err
}
func (r *MongoRepositoryTemplate) FindId(id interface{}, result interface{}, name string) error {
	defer r.connection.Close()
	err := r.getCollection(name).FindId(id).One(result)
	return err
}

func (r *MongoRepositoryTemplate) Insert(result interface{}, name string) error {
	defer r.connection.Close()
	err := r.getCollection(name).Insert(&result)
	return err
}

func (r *MongoRepositoryTemplate) Update(result interface{}, name string) error {
	defer r.connection.Close()
	domainAsserted := result.(domain.Domain)
	err := r.getCollection(name).UpdateId(domainAsserted.Id, result)
	return err
}

func (r *MongoRepositoryTemplate) Delete(result interface{}, name string) error {
	defer r.connection.Close()
	domainAsserted := result.(domain.Domain)
	err := r.getCollection(name).RemoveId(domainAsserted.Id)
	return err
}

func (r *MongoRepositoryTemplate) getCollection(name string) *mgo.Collection {
	session := r.connection.sessionCopy
	return session.DB(r.connection.database).C(name)
}

func (r *MongoRepositoryTemplate) SetConnection(connection MongoConnection) {
	r.connection = connection
}

type MongoConnection struct {
	sessionCopy *mgo.Session
	database    string
}

func (c *MongoConnection) Close() error {
	c.sessionCopy.Close()
	return nil
}

type MongoConnectionPool struct {
	mongoSession *mgo.Session
	database     string
}

func (p *MongoConnectionPool) GetConnection() MongoConnection {
	return MongoConnection{sessionCopy: p.mongoSession.Copy(), database: p.database}
}

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
