package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Conn
var Conn *sql.DB

//CreateSession creates the connection to database
func CreateSession() {

	Conn, err := sql.Open("mysql", "latest:latest@/latest")
	if err != nil {
		log.Println("Unable to connect to database")
	}

	log.Println(Conn.Ping)

}

//GetSession gets  an existing database connection

/*
func GetSession() *Conn {

	if Conn == nil {

		CreateSession()
	}

	return Conn

}  */

// Collection is for specific table
//func (db *DB) Collection(col string) mgo.Collection {
//return db.Database(db.dbName).Collection(col)

//}
