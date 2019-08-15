package model

import "gopkg.in/mgo.v2/bson"

type MgComment struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Ip       string        `bson:"ip"`
	Location string        `bson:"location"`
	Ts       int64         `bson:"ts"`
	Support  int           `bson:"support"`
	Oppose   int           `bson:"oppose"`
	Content  string        `bson:"content"`
	IsTop    bool          `bson:"istop"`

	Replys []bson.ObjectId `bson:"replys"`
	Topid  bson.ObjectId   `bson:"topid"`
	Refer  bson.ObjectId   `bson:"refer"`
}
