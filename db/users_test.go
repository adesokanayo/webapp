package db

// User define the data to be stored in the users db.
type User struct {
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
/*func Create(db *DB, NewUser interface{}) error {

	session := NewDB("test", Conn)

	defer session.Close()

	//c := session.dbName("test").C("user")
	c := session.DB("test").C("user")

	index := mgo.Index{
		Key:    []string{"id"},
		Unique: true,
	}

	s, err := c.Insert(NewUser)

	if err != nil {
		log.Println("User Creation failed")
	}

	return s
}

*/

//GetOneUser from the database

/*
func GetOneUser(db *mgo.Session, username string) {

	session := NewDB("test", db)

	defer session.Close()

	c := session.DB("test").C("user")

	index := mgo.Index{
		Key:    []string{"id"},
		Unique: true,
	}

	c.EnsureIndex(index)

	s := c.Find(&username)

	var result []User

	err := s.All(&result)

	if err != nil {
		log.Println("unable to show users")
	}

	return

}

//func Update(db *mgo.Session, userId string, user User) error{...}

//func Delete(db *mgo.Session, userId string){...}

//GetAllUsers is to retrieve all users on the system
func GetAllUsers(Conn *mgo.Session, username string) ([]User, error) {

	session := NewDB("test", Conn)

	defer session.Close()

	c := session.DB("test").C("user")

	index := mgo.Index{
		Key:    []string{"id"},
		Unique: true,
	}

	c.EnsureIndex(index)
	var result []User
	err := c.Find(nil).All(&result)

	if err != nil {
		log.Println("unable to show all users")
	}

	//return result, nil
	log.Println("Results:", result)
	return result, nil
}
*/
