package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type puppydb struct {
	Session *mongo.Client
}

func MongoConnect() (*puppydb, error) {
	// Set your MongoDB Atlas connection string here.
	connectionString := "mongodb+srv://aleatoryfreak:<password>@puppylove.woq42jd.mongodb.net/?retryWrites=true&w=majority"

	// Set client options.
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB Atlas cluster.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB Atlas: %v", err)
	}

	// Check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error pinging MongoDB Atlas: %v", err)
	}

	fmt.Println("Connected to MongoDB Atlas successfully!")

	return &puppydb{Session: client}, nil
}

func (db *puppydb) GetCollection(database, collection string) *mongo.Collection {
	return db.Session.Database(database).Collection(collection)
}

func (db *puppydb) FindById(database, collection string, id interface{}) (*mongo.SingleResult, error) {
	c := db.GetCollection(database, collection)
	return c.FindOne(context.Background(), id), nil
}
