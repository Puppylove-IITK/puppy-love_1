package main

import (
	"fmt"
	"os"

	"github.com/Puppylove-IITK/puppylove/config"
	"github.com/Puppylove-IITK/puppylove/router"
	"github.com/Puppylove-IITK/puppylove/utils"
	
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	config.CfgInit()

	mongoDb, error := db.MongoConnect()
	if error != nil {
		fmt.Print("[Error] Could not connect to MongoDB")
		fmt.Print("[Error] " + config.CfgMgoUrl)
		fmt.Print(os.Environ())
		os.Exit(1)
	}

	utils.Randinit()

	// set up session db
	store := cookie.NewStore([]byte(config.CfgAdminPass))

	// iris.Config.Gzip = true
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))
	router.PuppyRoute(r, mongoDb)
	if err := r.Run(config.CfgAddr); err != nil {
		fmt.Println("[Error] " + err.Error())
	}
}
