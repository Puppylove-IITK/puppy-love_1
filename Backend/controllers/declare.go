	package controllers

	import (
		"context"
		"log"
		"net/http"

		"go.mongodb.org/mongo-driver/mongo"
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

		cur, err := Db.GetCollection("user").Find(context.Background(), bson.M{"dirty": false}, options.Find())
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer cur.Close(context.Background())

		var people []typeIds
		if err := cur.All(context.Background(), &people); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var modelsList []mongo.WriteModel

		for _, pe := range people {
			res := models.NewDeclareTable(pe.Id)
			upsert := mongo.NewUpdateOneModel()
			upsert.SetFilter(res.Selector)
			upsert.SetUpdate(res.Change)
			upsert.SetUpsert(true)
			modelsList = append(modelsList, upsert)
		}

		bulkWriteOptions := options.BulkWrite().SetOrdered(false)
		r, err := Db.GetCollection("declare").BulkWrite(context.Background(), modelsList, bulkWriteOptions)

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, r)
	}
