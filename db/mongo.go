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

func Init() {
	Mongo.init()
}
func Destroy() {
	fmt.Println("destroying mongo")
	Mongo.session.Close()
}

type mongo struct {
	session  *mgo.Session
	database *mgo.Database
}

func (m *mongo) init() {
	fmt.Println("initiating mongo")
	session, err := mgo.Dial(dbconst.Uri)
	if err != nil {
		//TODO log
		fmt.Println(err)
		os.Exit(1)
	}
	m.session = session
	m.database = session.DB(dbconst.Database)
}

func (m *mongo) C(collection string) *mgo.Collection {
	return m.database.C(collection)
}
