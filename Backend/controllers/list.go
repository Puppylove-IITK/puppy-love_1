package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/Puppylove-IITK/puppylove/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// @AUTH Get user's basic information
// ---------------------------------------
type typeListAll struct {
	Id    string `json:"_id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Image string `json:"image" bson:"image"`
}

func ListAll(c *gin.Context) {
	var results []typeListAll

	// Fetch user
	cur, err := Db.GetCollection("user").Find(context.Background(), bson.M{})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Print(err)
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var result typeListAll
		err := cur.Decode(&result)
		if err != nil {
			log.Print(err)
			continue
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Print(err)
		return
	}

	c.JSON(http.StatusAccepted, results)
}

func PubkeyList(c *gin.Context) {
	var query []bson.M

	cursor, err := Db.GetCollection("user").
		Find(context.Background(), bson.M{"dirty": false})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Print(err)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user bson.M
		if err := cursor.Decode(&user); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Print(err)
			return
		}

		query = append(query, bson.M{
			"_id":    user["_id"],
			"pubKey": user["pubKey"],
		})
	}

	c.JSON(http.StatusAccepted, query)
}

func DeclareList(c *gin.Context) {
	id, err := SessionId(c)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	var resp models.Declare
	if err := Db.GetCollection("declare").FindOne(context.Background(), bson.M{"_id": id}).Decode(&resp); err != nil {
		if err == mongo.ErrNoDocuments {
			c.AbortWithStatus(http.StatusNotFound)
			log.Print(err)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	if resp.Token0 != "" {
		resp.Token0 = "1"
	}
	if resp.Token1 != "" {
		resp.Token1 = "1"
	}
	if resp.Token2 != "" {
		resp.Token2 = "1"
	}
	if resp.Token3 != "" {
		resp.Token3 = "1"
	}

	c.JSON(http.StatusOK, resp)
}
