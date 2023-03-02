package db

import (
	"crypto/tls"
	"net"

	"gopkg.in/mgo.v2"
)

type PuppyDb struct {
	S *mgo.Session
}

func MongoConnect() (PuppyDb, error) {

	var mongoURI = "mongodb+srv://aleatoryfreak:hFyRFQUC724RXS1q@puppylove.woq42jd.mongodb.net/?retryWrites=true&w=majority"

	dialInfo, err := mgo.ParseURL(mongoURI)

	//Below part is similar to above.
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	S, err := mgo.DialWithInfo(dialInfo)
	return PuppyDb{S}, err

	// ctx := context.TODO()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://aleatoryfreak:hFyRFQUC724RXS1q@puppylove.woq42jd.mongodb.net/?retryWrites=true&w=majority"))
	// return PuppyDb{Client: client}, err
}

func (db PuppyDb) GetById(table string, id string) *mgo.Query {
	return db.S.DB("puppy").C(table).FindId(id)
}

func (db PuppyDb) GetCollection(table string) *mgo.Collection {
	return db.S.DB("puppy").C(table)
}
