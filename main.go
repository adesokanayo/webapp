package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

var ses *mgo.Session

// GetSession exported so I can call it from users.go
func GetSession() *mgo.Session {

	if ses != nil {
		return ses.Copy()
	}

	session, err := mgo.Dial("localhost")
	if err != nil {

		log.Println("Unable to connect to server")
		panic(err)

	}

	log.Println("DB connection successful ")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	ses = session.Copy()

	return session
}

type myHandler struct {
}

func main() {

	connect := GetSession()

	fmt.Println(connect)

}
