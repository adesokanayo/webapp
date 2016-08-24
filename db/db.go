package main

import (

  //"log"
  //"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
  "fmt"


)

//DB is expopred type for database connection
type DB struct{
   *mgo.Session
   dbName string
}

//NewDB  is for every new connection.
func NewDB(db string, session *mgo.Session) *DB{
  return &DB{
      dbName: test,
      Session: session.Copy(),
   }
}

//Collection is for specific table
func (db *DB) Collection(col string) mgo.Collection{
  return db.Database(db.dbName).Collection(col)
}

func main() {

	connect := NewDB("localhost",DB)

	fmt.Println(connect)

}
