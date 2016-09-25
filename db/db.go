package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

//CreateSession creates the connection to database
func CreateSession() *mgo.Session {

	Conn, err := mgo.Dial("localhost")
	if err != nil {
		log.Println("Unable to connect to database")
	}
	Conn.SetMode(mgo.Monotonic, true)

	return Conn

}

//GetSession gets  an existing database connection
func GetSession() *mgo.Session {

	if mgoSession == nil {

		CreateSession()
	}

	return mgoSession

}

// Collection is for specific table
//func (db *DB) Collection(col string) mgo.Collection {
//return db.Database(db.dbName).Collection(col)

//}
