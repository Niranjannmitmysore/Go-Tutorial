// Sample program to show how to unmarshal a JSON document into
// a user defined struct type.
package main

import (
	"encoding/json"
	"fmt"
)

// document contains a JSON document.
var document = `{
"workshop": {
    "college_name": "Maharaja Institute of Technology Mysore",
    "company_name": "Qwinix"
},

"user":{
  "user_id": "4VV10IS045",
  "user_name": "praveen menon",
  "branch": "Information_science"
}
}`

// Fields to be encoded/decoded must be exported else the
// json encoding functions can't see the fields.

type (
  // UserDetails contains information about the user

  User_details struct {
    UserId string `json:"user_id"`
    UserName string `json:"user_name"`
    Branch string `json:"branch"`
  }

  // Workshop contains information for the Workshop for a User.

  Workshop struct {
		WorkshopDetails struct {
			CollegeName string `json:"college_name"`
      CompanyName string `json:"company_name"`
		} `json:"workshop"`
    UserDetails User_details `json:"user"`
	}
)

// main is the entry point for the application.
func main() {
	// Declare a variable of type Workshop.
	var uc Workshop

	// Unmarshal the JSON document into the variable.
	if err := json.Unmarshal([]byte(document), &uc); err != nil {
		fmt.Println(err)
		return
	}
  // Print the JSON
	fmt.Printf("%+v\n\n", uc)

  // Print required values from JSON

  fmt.Println("College Name:",uc.WorkshopDetails.CollegeName)
  fmt.Println("Company Name:",uc.WorkshopDetails.CompanyName)
  fmt.Println("User Name:",uc.UserDetails.UserName)
  fmt.Println("User Id:",uc.UserDetails.UserId)

}
