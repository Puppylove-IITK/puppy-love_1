package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Puppylove-IITK/puppylove/config"
	"github.com/Puppylove-IITK/puppylove/db"
	"github.com/Puppylove-IITK/puppylove/router"
	"github.com/Puppylove-IITK/puppylove/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// config.CfgInit()
	Db, err := db.MongoConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Session.Disconnect(context.Background())

	utils.Randinit()

	// set up session db
	store := cookie.NewStore([]byte(config.CfgAdminPass))

	// iris.Config.Gzip = true
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))
	router.PuppyRoute(r, *Db)
	if err := r.Run(config.CfgAddr); err != nil {
		fmt.Println("[Error] " + err.Error())
	}
}
