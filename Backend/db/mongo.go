package db

import (
	"crypto/tls"
	"net"
	"time"

	"gopkg.in/mgo.v2"
)

type PuppyDb struct {
	S *mgo.Session
}

func MongoConnect() (PuppyDb, error) {
	// MongoDB Atlas connection URI
	uri := "mongodb+srv://aleatoryfreak:hFyRFQUC724RXS1q@puppylove.woq42jd.mongodb.net/?retryWrites=true&w=majority"

	dialInfo, err := mgo.ParseURL(uri)
	if err != nil {
		return PuppyDb{}, err
	}

	// Configure TLS
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		if err != nil {
			return nil, err
		}
		return conn, nil
	}

	// Set timeouts
	dialInfo.Timeout = 10 * time.Second

	S, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return PuppyDb{}, err
	}

	// Set safe and mode
	S.SetSafe(&mgo.Safe{})
	S.SetMode(mgo.Monotonic, true)

	return PuppyDb{S}, nil
}

func (db PuppyDb) GetById(table string, id string) *mgo.Query {
	return db.S.DB("puppy").C(table).FindId(id)
}

func (db PuppyDb) GetCollection(table string) *mgo.Collection {
	return db.S.DB("puppy").C(table)
}
