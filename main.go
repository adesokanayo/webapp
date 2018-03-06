package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	Router := mux.NewRouter().StrictSlash(false)
	Router.HandleFunc("/api/test", Test)
	/*
		Router.HandleFunc("/api/users/{username}", GetUser).Methods("GET")
		Router.HandleFunc("/api/users", GetAllUsers).Methods("GET")
		Router.HandleFunc("/api/users", CreateUser).Methods("POST")
		Router.HandleFunc("/api/users/{username}", DeleteUser).Methods("DELETE")
	*/
	server := &http.Server{
		Addr:    ":3000",
		Handler: Router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
	//db.CreateSession()

}

//Test Handler: I used it to test my server
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>let's go there </h1>")

	Conn, err := sql.Open("mysql", "latest:latest@/latest")
	if err != nil {
		panic(err)
	}
	fmt.Println("COnnection Open Succesfully")

	log.Println(Conn.Ping)
}

/*
//GetUser  is to get one single user
func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	URLParams := mux.Vars(r)
	username := URLParams["username"]
	fmt.Println(username)

	result := db.User{}
	newsession := db.CreateSession()

	defer newsession.Close()

	c := newsession.DB("test").C("users")

	index := mgo.Index{
		Key:    []string{"id"},
		Unique: true,
	}

	c.EnsureIndex(index)

	err := c.Find(bson.M{"username": username}).One(&result)

	if err != nil {
		log.Println("User not Found")
		fmt.Println(err)
	}

	y, _ := json.Marshal(&result)
	if err != nil {
		log.Println("unable to   interpret the output")
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(y)

}

//GetAllUsers is handler to return all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	result := []db.User{}

	newsession := db.CreateSession()

	defer newsession.Close()

	c := newsession.DB("test").C("users")

	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		fmt.Println("Quey failed to execute ")
	}

	y, err := json.Marshal(&result)
	if err != nil {
		log.Println("unable to call the interpret the output")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(y)

}

//CreateUser is to create new user on the platform
func CreateUser(w http.ResponseWriter, r *http.Request) {

	s := db.CreateSession()
	defer s.Close()
	ru := db.NewUser{}

	ru.Username = r.FormValue("username")
	ru.Password = r.FormValue("password")
	ru.FirstName = r.FormValue("firstname")
	ru.LastName = r.FormValue("lastname")

	var newcomer db.NewUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newcomer)
	fmt.Println(newcomer)

	if err != nil {
		fmt.Println(err)
		return
	}

	c := s.DB("test").C("users")

	fmt.Println(newcomer)

	err = c.Insert(newcomer)

	if err != nil {
		fmt.Println("unable to insert into db")
	}

	log.Println("Successfully Inserted into Db")
	w.WriteHeader(http.StatusCreated)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	newsession := db.CreateSession()
	defer newsession.Close()

	w.Header().Set("Content-Type", "application/json")

	URLParams := mux.Vars(r)
	username := URLParams["username"]
	fmt.Println(username)

	c := newsession.DB("test").C("users")

	err := c.Remove(bson.M{"username": username})

	if err != nil {
		log.Println("unable to delete users")
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)

}

*/
