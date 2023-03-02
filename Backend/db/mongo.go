package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PuppyDb struct {
	Client *mongo.Client
}

func MongoConnect() (PuppyDb, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://aleatoryfreak:hFyRFQUC724RXS1q@puppylove.woq42jd.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), clientOptions)
	return PuppyDb{Client: client}, err
}

func (db PuppyDb) GetById(table string, id string) *mongo.SingleResult {
	collection := db.Client.Database("puppy").Collection(table)
	filter := bson.M{"_id": id}
	return collection.FindOne(context.Background(), filter)
}

func (db PuppyDb) GetCollection(table string) *mongo.Collection {
	return db.Client.Database("puppy").Collection(table)
}
