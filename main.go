package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adesokanayo/webapp/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	//http.ListenAndServe(":1212", http.FileServer(http.Dir("public")))

	serve()

	//http.ListenAndServe(":1212", http.FileServer(http.Dir("public")))

	//Router.HandleFunc("/api/users/{username}", GetUser).Methods("GET")
	//Router.HandleFunc("/api/users", GetAllUsers).Methods("GET")
	//Router.HandleFunc("/api/users", CreateUser).Methods("POST")
	//Router.HandleFunc("/api/users/{username}", DeleteUser).Methods("DELETE")

}

func serve() {

	port := os.Getenv("port")

	if port == "" {
		port = "1212"
	}

	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	//http.Handle("/images/", http.FileServer(http.Dir("public/images")))
	//http.Handle("/css/", http.FileServer(http.Dir("public/css")))
	//http.Handle("/js/", http.FileServer(http.Dir("public/js")))

	http.HandleFunc("/public/", handlers.ServeResource)
	http.HandleFunc("/images/", handlers.ServeResource)
	http.HandleFunc("/css/", handlers.ServeResource)
	http.HandleFunc("/js/", handlers.ServeResource)
	http.HandleFunc("/assets/", handlers.ServeResource)
	http.HandleFunc("/assets/css/", handlers.ServeResource)
	http.HandleFunc("/assets/images/", handlers.ServeResource)
	http.HandleFunc("/starters/", handlers.ServeResource)
	http.HandleFunc("/starters/assets/", handlers.ServeResource)
	http.HandleFunc("/starters/assetsimages/", handlers.ServeResource)
	http.HandleFunc("/starters/assets/css/", handlers.ServeResource)
	http.HandleFunc("/temp/", handlers.ServeResource)

	//http.HandleFunc("/starters/", handlers.ServeResource)

	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/index", handlers.Index)
	http.HandleFunc("/register", handlers.Register)

	Router := mux.NewRouter().StrictSlash(false)
	//Router.HandleFunc("/api/test", Test)
	http.Handle("/", Router)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("server error: there was a problem", nil)
	} else {
		fmt.Println("Now serving with serve method")
	}

}

//Test Handler: I used it to test my server
/*
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>let's go there </h1>")

	Conn, err := sql.Open("mysql", "latest:latest@/latest")
	if err != nil {
		panic(err)
	}
	fmt.Println("COnnection Open Succesfully")

	log.Println(Conn.Ping)
}

*/

/*

