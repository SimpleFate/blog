package services

import (
	"blog/db"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

var (
	commentCol = db.Mongo.C("comment")
)

/*
留言板
*/

func SupportComment(id string) {
	_ = commentCol.Update(
		bson.M{
			"_id": bson.ObjectId(id),
		},
		bson.M{
			"$inc": bson.M{
				"support": 1,
			},
		},
	)
}

func OpposeComment(id string) {
	_ = commentCol.Update(
		bson.M{
			"_id": bson.ObjectId(id),
		},
		bson.M{
			"$inc": bson.M{
				"oppose": 1,
			},
		},
	)
}

func AddComment(privacy *Privacy, name, content string) {
	if privacy == nil {
		privacy = &Privacy{
			Ip:       "",
			Location: "",
			Dt:       0,
		}
	}
	if name == "" || content == "" {
		return
	}

	err := commentCol.Insert(bson.M{
		"name":     name,
		"content":  content,
		"ip":       privacy.Ip,
		"location": privacy.Location,
		"ts":       privacy.Dt,
		"support":  0,
		"oppose":   0,
		"istop":    true,

		"replys": make([]bson.ObjectId, 0),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
