package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Basic is okay
type Basic struct {
	fname string
	Lname string
	Email string
	Sex   string
}

type Secondary struct {
	Bio    string
	DOB    string
	Joined string
}

type WorkInfo struct {
	Comp     string
	Role     string
	CompAddr string
	Email    string
}

const (
	dbtype    = "mysql"
	dbpath    = "latest:latest@/latest"
	userexist = "SELECT COUNT(*) FROM users WHERE user = ?"
	signup    = "INSERT INTO users (nick, fname, sname, email,phone,bio,rel,dob,sex,joined,password) VALUES(?,?,?,?,?,?,?,?,?,?,?)"
)

var (
	// DBCon is the connection handle for the database
	DBCon *sql.DB
)

//OpenConn is for
func OpenConn() *sql.DB {

	db, err := sql.Open(dbtype, dbpath)

	if err != nil {
		fmt.Println("unable to open conn")
	} else {
		fmt.Println("connected to database by calling the OpenConn Function")
	}

	fmt.Println(db.Ping())

	return db

}
