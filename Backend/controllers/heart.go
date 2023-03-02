package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DbName = "puppy"
)

type Database struct {
	Client *mongo.Client
}

func (db *Database) Connect() error {
	var err error

	// Create a client instance
	db.Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://aleatoryfreak:hFyRFQUC724RXS1q@puppylove.woq42jd.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = db.Client.Connect(ctx)
	if err != nil {
		return err
	}

	return nil
}

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

	// Create a database instance
	var db Database
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	defer db.Client.Disconnect(context.TODO())

	// Access the database
	database := db.Client.Database(DbName)

	// Fetch user
	// filter := bson.D{
	// 	{"time", bson.D{
	// 		{"$gt", ltime},
	// 		{"$lte", ctime},
	// 	}},
	// }

	cur, err := database.Collection("heart").Find(context.TODO(), bson.M{"time": bson.M{"$gt": ltime, "$lte": ctime}})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Print(err)
		return
	}
	defer cur.Close(context.Background())

	if err := cur.All(context.Background(), &votes); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Print(err)
		return
	}

	c.JSON(http.StatusAccepted, bson.M{
		"votes": votes,
		"time":  ctime,
	})
}
