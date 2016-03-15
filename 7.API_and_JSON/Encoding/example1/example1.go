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
}
}`

// Fields to be encoded/decoded must be exported else the
// json encoding functions can't see the fields.

type (
	// Workshop contains information for the user.
	Workshop struct {
		WorkshopDetails struct {
			CollegeName string `json:"college_name"`
			CompanyName string `json:"company_name"`
			} `json:"workshop"`
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
		fmt.Printf("%+v\n\n", uc)
		fmt.Println("CollegeName:",uc.WorkshopDetails.CollegeName)
		fmt.Println("CompanyName:",uc.WorkshopDetails.CompanyName)

	}
