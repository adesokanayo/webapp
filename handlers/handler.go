package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//NewUser has basic
type NewUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Status    int    `json:"status"`
}

//ServeResource is good
func ServeResource(w http.ResponseWriter, r *http.Request) {
	path := "public/" + r.URL.Path
	http.ServeFile(w, r, path)
}

//Index is alright
func Index(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("Public/index.html")
	if err != nil {
		fmt.Println("someting bad happend")
	}
	t.Execute(w, nil)
	fmt.Println("I called index successfully")

}

//Login is for login
func Login(w http.ResponseWriter, r *http.Request) {

	log.Println("now listening")
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Public/login_background.html")
		t.Execute(w, nil)
	}

	//var N NewUser

	///err := dbconn.QueryRow("select username from users where username= ?", "ade").Scan(&N.Username)
	///if err != nil {
	///	fmt.Println("somthing happened", error(err))
	//}

	//if err ! = nil { fmt.Println("unable to execute query ",error(err))

	//fmt.Println(prepstat)

}

//Register is okay
func Register(w http.ResponseWriter, r *http.Request) {

	log.Println("now listenting for register")
	if r.Method == "GET" {
		t, err := template.ParseFiles("public/login_registration.html")
		if err != nil {
			panic(err)

		}
		t.Execute(w, nil)
	}

	//t, err := template.ParseFiles("public/index.html")
	//if err != nil {
	//	log.Println("within here mm", error(err))

	//	t.Execute(w, nil)

	//}

}
