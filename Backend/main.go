package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Puppylove-IITK/puppylove/db"
)

func main() {
	// config.CfgInit()
	Db, err := db.MongoConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Session.Disconnect(context.Background())

	id := "1234"
	result, err := db.FindById("puppy", "dogs", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	collection := db.GetCollection("puppy", "dogs")
	fmt.Println(collection)

	// mongoDb, error := db.MongoConnect()
	// if error != nil {
	// 	fmt.Print("[Error] Could not connect to MongoDB")
	// 	fmt.Print("[Error] " + config.CfgMgoUrl)
	// 	fmt.Print(os.Environ())
	// 	os.Exit(1)
	// }

	// utils.Randinit()

	// // set up session db
	// store := cookie.NewStore([]byte(config.CfgAdminPass))

	// iris.Config.Gzip = true
	// r := gin.Default()
	// r.Use(sessions.Sessions("mysession", store))
	// router.PuppyRoute(r, mongoDb)
	// if err := r.Run(config.CfgAddr); err != nil {
	// 	fmt.Println("[Error] " + err.Error())
	// }
}
