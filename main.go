package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adesokanayo/webapp/db"
	"github.com/gorilla/mux"
)

func main() {

	db.CreateSession()

	Router := mux.NewRouter().StrictSlash(false)
	Router.HandleFunc("/api/users/", GetAllUserHandler)
	Router.HandleFunc("/api/notes", Test)
	Router.HandleFunc("/api/user/{username}", GetUserHandler).Methods("GET")
	//r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	//r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	server := &http.Server{
		Addr:    ":3000",
		Handler: Router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}

//Test Handler: I used it to test my server
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>let's go there </h1>")

}

//GetUserHandler  is to get one single user
func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//database := db.CreateSession()

	URLParams := mux.Vars(r)
	username := URLParams["username"]
	fmt.Println(username)

	err := db.GetOneUser(db.GetSession(), username)
	if err != nil {
		log.Println("unable to call the Get One User method")
	}

	log.Println("called the single handler ")
	y, err := json.Marshal(&username)
	if err != nil {
		log.Println("unable to Marshal the interpret the output")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(y)

}

//GetAllUserHandler is handler for all users
func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("called the service to get all users ")
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	log.Println(r.URL.Path)
	database := db.GetSession()
	j, err := db.GetAllUsers(database, "adesokano")

	if err != nil {
		log.Println("unable to call the GetAllUsers method")
	}

	y, err := json.Marshal(j)
	if err != nil {
		log.Println("unable to call the interpret the output")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(y)

}
