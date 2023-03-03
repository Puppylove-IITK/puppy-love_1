package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/Puppylove-IITK/puppylove/db"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Puppylove-IITK/puppylove/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// @AUTH @Admin Create the entries in the declare table
// ----------------------------------------------------
func DeclarePrepare(c *gin.Context) {
	id, err := SessionId(c)
	if err != nil || id != "admin" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	type typeIds struct {
		Id string `json:"_id" bson:"_id"`
	}

	cur, err := db.GetCollection("user").Find(context.Background(), bson.M{"dirty": false}, options.Find())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var people []typeIds

	collection := db.Client().Database("test").Collection("user")
	filter := bson.M{"dirty": false}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var pe typeIds
		if err := cur.Decode(&pe); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		people = append(people, pe)
	}
	if err := cur.Err(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	declareCollection := db.Client().Database("test").Collection("declare")
	var models []mongo.WriteModel
	for _, pe := range people {
		res := models.NewDeclareTable(pe.Id)
		models = append(models, res)
	}

	opts := options.BulkWrite().SetOrdered(false)
	r, err := declareCollection.BulkWrite(context.Background(), models, opts)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, r)
}

