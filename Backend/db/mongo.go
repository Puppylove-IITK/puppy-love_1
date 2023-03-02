package db

import (
	"github.com/Puppylove-IITK/puppylove/config"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PuppyDb struct {
	S *mgo.Session
}

func MongoConnect() (PuppyDb, error) {
	// Create MongoDB DialInfo object using MongoDB URI with authentication credentials
	mongoInfo := &mgo.DialInfo{
		Addrs:    []string{config.CfgMgoUrl},
		Database: config.CfgMgoDb,
		Username: config.CfgMgoUser,
		Password: config.CfgMgoPass,
	}

	// Create a new session using DialWithInfo
	S, err := mgo.DialWithInfo(mongoInfo)
	return PuppyDb{S}, err
}

func (db PuppyDb) GetById(table string, id string) *mgo.Query {
	return db.S.DB(config.CfgMgoDb).C(table).FindId(bson.ObjectIdHex(id))
}

func (db PuppyDb) GetCollection(table string) *mgo.Collection {
	return db.S.DB(config.CfgMgoDb).C(table)
}
