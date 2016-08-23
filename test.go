package main

import "fmt"

type response1 struct {
	Page   int
	Fruits string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	y := response1{4, "Apple"}

	fmt.Printf("Yes:", y)
}
