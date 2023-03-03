package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect() {
	// Set your MongoDB Atlas connection string here.
	connectionString := "mongodb+srv://aleatoryfreak:<password>@puppylove.woq42jd.mongodb.net/?retryWrites=true&w=majority"

	// Set client options.
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB Atlas cluster.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB Atlas:", err)
		return
	}

	// Check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB Atlas:", err)
		return
	}

	fmt.Println("Connected to MongoDB Atlas successfully!")

	return
}
