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
