package users

import (

	"log"
	"db"

)

// Users define the data to be stored in the users db.
type Users struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Status    int    `json:"status"`
}

// NewUser defines the data recieved from the front end to register a new user
type NewUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Status    int    `json:"status"`
}



//Create  is also an exported function to create new users.
func Create(db *DB, user NewUser) error {

session=newDB()

defer session.Close()



	c := session.db("test").C("user")

	index := mgo.Index{
		Key:    []string{"id"},
		Unique: true,
	}

	err, s := c.Insert(&NewUser)

	if err != nil {
		log.Println("User Creation failed")
	}

	return s
}

//func Get(db *mgo.Session, userId string) (*User, error) {...}

//func Update(db *mgo.Session, userId string, user User) error{...}

//func Delete(db *mgo.Session, userId string){...}
