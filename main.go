package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adesokanayo/webapp/db"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	Router := mux.NewRouter().StrictSlash(false)
	Router.HandleFunc("/api/test", Test)
	Router.HandleFunc("/api/users/{username}", GetUser).Methods("GET")
	Router.HandleFunc("/api/users", GetAllUsers).Methods("GET")
	Router.HandleFunc("/api/users", CreateUser).Methods("POST")
	//Router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
	server := &http.Server{
		Addr:    ":3000",
		Handler: Router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
	db.CreateSession()
}

//Test Handler: I used it to test my server
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>let's go there </h1>")

}

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
		log.Println("unable to show users")
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
		fmt.Println("something gone bad ")
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

	//var newcomer db.NewUser

}
