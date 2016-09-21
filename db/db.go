package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

//DB is exported type for database connection
type DB struct {
	*mgo.Session
	dbName string
}

// Conn is exported
var Conn *mgo.Session

//NewDB returns a new database .
func NewDB(db string, session *mgo.Session) *DB {
	return &DB{
		dbName:  db,
		Session: session.Copy(),
	}
}

//CreateSession is expopred
func CreateSession() {

	session, err := mgo.Dial("localhost")
	if err != nil {

		log.Println("Unable to connect to database")
		//panic(err)
	}

	log.Println("DB connection successful ")

	session.SetMode(mgo.Monotonic, true)

	defer session.Close()

}

func GetSession() *mgo.Session {

	return Conn
}

// Collection is for specific table
//func (db *DB) Collection(col string) mgo.Collection {
//	return db.Database(db.dbName).Collection(col)
//}
