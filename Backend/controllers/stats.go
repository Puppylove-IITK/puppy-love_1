package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/puppy-love/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetStats returns useful statistics
func GetStats(c *gin.Context) {
	collectionUsers := Db.GetCollection("user")
	collectionHearts := Db.GetCollection("heart")

	var users []models.User
	var hearts []models.Heart

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filterUsers := bson.M{"dirty": false}
	filterHearts := bson.M{}

	if cur, err := collectionUsers.Find(ctx, filterUsers); err != nil {
		c.String(http.StatusInternalServerError, "Could not get database info")
		return
	} else {
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var user models.User
			if err := cur.Decode(&user); err != nil {
				c.String(http.StatusInternalServerError, "Could not get database info")
				return
			}
			users = append(users, user)
		}
	}

	if cur, err := collectionHearts.Find(ctx, filterHearts); err != nil {
		c.String(http.StatusInternalServerError, "Could not get database info")
		return
	} else {
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var heart models.Heart
			if err := cur.Decode(&heart); err != nil {
				c.String(http.StatusInternalServerError, "Could not get database info")
				return
			}
			hearts = append(hearts, heart)
		}
	}

	var y22males, y21males, y20males, y19males, y18males, othermales int
	var y22females, y21females, y20females, y19females, y18females, otherfemales int

	for _, user := range users {
		if user.Gender == "1" {
			if strings.HasPrefix(user.Id, "22") {
				y22males++
			} else if strings.HasPrefix(user.Id, "21") {
				y21males++
			} else if strings.HasPrefix(user.Id, "20") {
				y20males++
			} else if strings.HasPrefix(user.Id, "19") {
				y19males++
			} else if strings.HasPrefix(user.Id, "18") {
				y18males++
			} else {
				othermales++
			}
		} else {
			if strings.HasPrefix(user.Id, "22") {
				y22females++
			} else if strings.HasPrefix(user.Id, "21") {
				y21females++
			} else if strings.HasPrefix(user.Id, "20") {
				y20females++
			} else if strings.HasPrefix(user.Id, "19") {
				y19females++
			} else if strings.HasPrefix(user.Id, "18") {
				y18females++
			} else {
				otherfemales++
			}
		}
	}

	var y22maleHearts, y21maleHearts, y20maleHearts, y19maleHearts, y18maleHearts, othermaleHearts int
	var y22femaleHearts, y21femaleHearts, y20femaleHearts, y19femaleHearts, y18femaleHearts, otherfemaleHearts int

	for _, heart := range hearts {
		if heart.Gender == "1" {
			if strings.HasPrefix(heart.Id, "22") {
				y22maleHearts++
			} else if strings.HasPrefix(heart.Id, "21") {
				y21maleHearts++
			} else if strings.HasPrefix(heart.Id, "20") {
				y20maleHearts++
			} else if strings.HasPrefix(heart.Id, "19") {
				y19maleHearts++
			} else if strings.HasPrefix(heart.Id, "18") {
				y18maleHearts++
			} else {
				othermaleHearts++
			}
		} else {
			if strings.HasPrefix(heart.Id, "22") {
				y22femaleHearts++
			} else if strings.HasPrefix(heart.Id, "21") {
				y21femaleHearts++
			} else if strings.HasPrefix(heart.Id, "20") {
				y20femaleHearts++
			} else if strings.HasPrefix(heart.Id, "19") {
				y19femaleHearts++
			} else if strings.HasPrefix(heart.Id, "18") {
				y18femaleHearts++
			} else {
				otherfemaleHearts++
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"users":             len(users),
		"y22males":          y22males,
		"y21males":          y21males,
		"y20males":          y20males,
		"y19males":          y19males,
		"y18males":          y18males,
		"othermales":        othermales,
		"y22females":        y22females,
		"y21females":        y21females,
		"y20females":        y20females,
		"y19females":        y19females,
		"y18females":        y18females,
		"otherfemales":      otherfemales,
		"y22maleHearts":     y22maleHearts,
		"y21maleHearts":     y21maleHearts,
		"y20maleHearts":     y20maleHearts,
		"y19maleHearts":     y19maleHearts,
		"y18maleHearts":     y18maleHearts,
		"othermaleHearts":   othermaleHearts,
		"y22femaleHearts":   y22femaleHearts,
		"y21femaleHearts":   y21femaleHearts,
		"y20femaleHearts":   y20femaleHearts,
		"y19femaleHearts":   y19femaleHearts,
		"y18femaleHearts":   y18femaleHearts,
		"otherfemaleHearts": otherfemaleHearts,
	})
}
