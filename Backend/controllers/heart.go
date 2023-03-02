package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Puppylove-IITK/puppylove/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HeartGet(c *gin.Context) {
	id, err := SessionId(c)
	if err != nil || id != c.Param("you") {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	// Last checked time
	ltime, err := strconv.ParseUint(c.Param("time"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad timestamp value")
		return
	}

	// Current time
	ctime := uint64(time.Now().UnixNano() / 1000000)

	type AnonymVote struct {
		Value          string `json:"v" bson:"v"`
		GenderOfSender string `json:"genderOfSender" bson:"gender"`
	}

	votes := new([]AnonymVote)

	// Create a client instance
	client, err := mongo.NewClient(options.Client().ApplyURI(DbUrl))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Access the database
	db := client.Database(DbName)

	// Fetch user
	if err := db.Collection("heart").
		Find(context.TODO(), bson.M{"time": bson.M{"$gt": ltime, "$lte": ctime}}).
		All(context.Background(), votes); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Print(err)
		return
	}

	if *votes == nil {
		*votes = []AnonymVote{}
	}

	c.JSON(http.StatusAccepted, bson.M{
		"votes": *votes,
		"time":  ctime,
	})
}
