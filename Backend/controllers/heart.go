package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	// TODO: fix bindings to be consistent across dbs
	type AnonymVote struct {
		Value          string `json:"v" bson:"v"`
		GenderOfSender string `json:"genderOfSender" bson:"gender"`
	}

	var votes []AnonymVote

	// Fetch votes
	cursor, err := Db.GetCollection("heart").Find(context.Background(), bson.M{"time": bson.M{"$gt": ltime, "$lte": ctime}})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Print(err)
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &votes); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Print(err)
		return
	}

	if len(votes) == 0 {
		votes = []AnonymVote{}
	}

	c.JSON(http.StatusAccepted, bson.M{
		"votes": votes,
		"time":  ctime,
	})
}
