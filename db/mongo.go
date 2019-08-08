package db

import (
	"blog/consts/dbconst"
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

var (
	Mongo = &mongo{}
)

func init() {
	Mongo.init()
}

type mongo struct {
	session  *mgo.Session
	database *mgo.Database
}

func (m *mongo) init() {
	session, err := mgo.Dial(dbconst.Uri)
	if err != nil {
		//TODO log
		fmt.Println(err)
		os.Exit(1)
	}
	m.session = session
	m.database = session.DB(dbconst.Database)
	fmt.Println("initiating mongo")

}

func (m *mongo) C(collection string) *mgo.Collection {
	return m.database.C(collection)
}
