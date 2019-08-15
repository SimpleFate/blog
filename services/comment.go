package services

import (
	"blog/db"
	"blog/model"
	"blog/utils"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
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
			"_id": bson.ObjectIdHex(id),
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
			"_id": bson.ObjectIdHex(id),
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

func ReplyComment(originId string, name string, remark string, privacy *Privacy) {
	origin := model.MgComment{}
	replyId := bson.NewObjectId()
	reply := bson.M{
		"_id":      replyId,
		"name":     name,
		"content":  remark,
		"ip":       privacy.Ip,
		"location": privacy.Location,
		"ts":       privacy.Dt,
		"support":  0,
		"oppose":   0,
		"istop":    false,
	}

	err := commentCol.Find(bson.M{
		"_id": bson.ObjectIdHex(originId),
	}).One(&origin)
	if err != nil {
		return
	}
	if origin.IsTop {
		reply["topid"] = origin.Id
		reply["refer"] = origin.Id
		_ = commentCol.Insert(&reply)
		_ = commentCol.Update(bson.M{
			"_id": origin.Id,
		},
			bson.M{
				"$push": bson.M{
					"replys": replyId,
				},
			})
	} else {
		reply["refer"] = origin.Id
		reply["topid"] = origin.Topid
		_ = commentCol.Insert(reply)
		_ = commentCol.Update(bson.M{
			"_id": origin.Topid,
		},
			bson.M{
				"$push": bson.M{
					"replys": replyId,
				},
			})
	}

}

//id, name, content, time, ip, address, remarks, support, oppose
func ListTopComments() []map[string]string {
	query := commentCol.Find(bson.M{
		"istop": true,
	}).Sort("ts")
	iter := query.Iter()
	count, err := query.Count()
	if err != nil {
		count = 0
	}
	res := make([]map[string]string, 0, count)
	comment := model.MgComment{}
	for iter.Next(&comment) {
		item := make(map[string]string)
		item["id"] = comment.Id.Hex()
		item["name"] = comment.Name
		item["content"] = comment.Content
		item["ip"] = comment.Ip
		item["address"] = comment.Location
		item["time"] = utils.DtToString(comment.Ts)
		item["remarks"] = strconv.Itoa(len(comment.Replys))
		item["support"] = strconv.Itoa(comment.Support)
		item["oppose"] = strconv.Itoa(comment.Oppose)
		res = append(res, item)
	}
	return res
}

//id, name, content, time, ip, address, support, oppose, referName 如果回复顶层评论就为空字符串,
func GetReplys(commentId string) []map[string]string {
	query := commentCol.Find(bson.M{
		"topid": bson.ObjectIdHex(commentId),
	})
	iter := query.Iter()
	count, err := query.Count()
	if err != nil {
		count = 0
	}
	res := make([]map[string]string, 0, count)
	comment := model.MgComment{}
	for iter.Next(&comment) {
		item := make(map[string]string)
		item["id"] = comment.Id.Hex()
		item["name"] = comment.Name
		item["content"] = comment.Content
		item["ip"] = comment.Ip
		item["address"] = comment.Location
		item["time"] = utils.DtToString(comment.Ts)
		item["support"] = strconv.Itoa(comment.Support)
		item["oppose"] = strconv.Itoa(comment.Oppose)
		referName := ""
		if comment.Refer != bson.ObjectIdHex(commentId) {
			referRes := struct {
				Name string `bson:"name"`
			}{}
			err = commentCol.Find(bson.M{
				"_id": comment.Refer,
			}).Select(bson.M{
				"name": 1,
			}).One(&referRes)
			if err == nil {
				referName = referRes.Name
			}
		}
		item["referName"] = referName
		res = append(res, item)
	}
	return res
}
