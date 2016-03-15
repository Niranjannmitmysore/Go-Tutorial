
// Sample program to show how to unmarshal a JSON request from Postman and send JSON response
// a user defined struct type from a file.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
)


type (

	// User struct contains information for the user.
	User struct {
		Name string
		Email string
		Address string
	}

	Response struct {
		Name string
		Email string
		Address string
		Status string
		Message string
	}
)

//Create Function accepts the input to Create a user

func Create(rw http.ResponseWriter, req *http.Request) {

	// Read all data from request BODY and copy it to a variable named body
	JSON_body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	// Declare a variable of type User.
	var u User

	// Unmarshal the JSON_body into the variable.
	err = json.Unmarshal(JSON_body, &u)
	if err != nil {
		panic(err)
	}

	// Print data from User struct
	fmt.Println("\n Name:",u.Name + "\n Email:",u.Email + "\n Address:",u.Address)

	// Marshal the JSON_body and create a response

	result, err := json.Marshal(Response{
		Name:     u.Name,
		Email:    u.Email,
		Address:  u.Address,
		Status:   "Success",
		Message:  "User is successfully signed in",
	})

	if err != nil {
		log.Fatal(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(result)
}

// main is the entry point for the application.
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sign_up", Create).Methods("POST")

	// http.Handle("/", r)
	handler := cors.Default().Handler(r)
	http.Handle("/", handler)

	// HTTP Listening Port
	log.Println("main : Started : Listening on: http://localhost:3000 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))

}
