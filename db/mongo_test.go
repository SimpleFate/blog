package db

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestMongo_C(t *testing.T) {
	iter := Mongo.C("admin").Find(bson.M{}).Iter()
	defer iter.Close()
	res := bson.M{}
	for iter.Next(res) {
		fmt.Println("finded : ", res)
	}
}
