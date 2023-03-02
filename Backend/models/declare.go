import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Declare struct {
		Id     string `json:"_id" bson:"_id"`
		Token0 string `json:"t0" bson:"t0"`
		Token1 string `json:"t1" bson:"t1"`
		Token2 string `json:"t2" bson:"t2"`
		Token3 string `json:"t3" bson:"t3"`
	}
	PairUpsert struct {
		Selector bson.M
		Change   bson.M
	}
)

// Create table update object for token table
func UpsertDeclareTable(d *Declare) *options.UpdateOptions {
	return options.Update().SetUpsert(true)
}

func NewDeclareTable(id string) *options.FindOneAndUpdateOptions {
	selector := bson.M{"_id": id}
	change := bson.M{"$setOnInsert": bson.M{
		"_id": id,
		"t0":  "",
		"t1":  "",
		"t2":  "",
		"t3":  "",
	}}
	return options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
}
