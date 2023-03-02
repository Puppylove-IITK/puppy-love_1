package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/Puppylove-IITK/puppylove/models"
	"github.com/Puppylove-IITK/puppylove/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	var people []typeIds

	cur, err := Db.GetCollection("user").Find(context.Background(), bson.M{"dirty": false})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for cur.Next(context.Background()) {
		var pe typeIds
		err := cur.Decode(&pe)
		if err != nil {
			log.Println(err)
			continue
		}
		res := models.NewDeclareTable(pe.Id)
		opt := options.Update().SetUpsert(true)
		if _, err := Db.GetCollection("declare").UpdateOne(context.Background(), res.Selector, res.Change, opt); err != nil {
			log.Println(err)
		}
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
	}

	c.Status(http.StatusOK)
}
