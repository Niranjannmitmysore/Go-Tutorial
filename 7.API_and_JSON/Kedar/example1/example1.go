package main

import (
"encoding/json"
"fmt"
)

/* Go offers built in support for JSON encoding and  decoding, including to and from built-in and custom data types. */

/* Let's define 2 structs */

type Vehicle1 struct {
	Type int
	Cars []string
}

type Vehicle2 struct {
	Type int `json:"type"`
	Cars []string `json:"cars"`
}

func main() {
/* We will encode slices and maps into JSON array and objects. */
	slice1 := []string{"BMW", "Mercedes", "Range Rover", "Hummer"}
	marshal_slice1, err := json.Marshal(slice1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal_slice1))	

	map1 := map[string]int{"BMW": 1, "Mercedes": 2, "Range Rover": 3, "Hummer": 4}
	marshal_map1, err := json.Marshal(map1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal_map1))

	/* The json package can automatically encode custom data types. It will include exported fields in the 
encoded output and will by default use those names as the keys. */
	response1 := &Vehicle1{
		Type: 1,
		Cars: []string{"BMW", "Mercedes", "Range Rover", "Hummer"}}
		marshal_response1, err := json.Marshal(response1)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(marshal_response1))

	/* We can also make use of tags on struct field declaration to customize encoded keys. */
		response2 := &Vehicle2{
			Type: 1,
			Cars: []string{"BMW", "Mercedes", "Range Rover", "Hummer"}}
			marshal_response2, err := json.Marshal(response2)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(marshal_response2))

	/* Let's understand how to decode JSON data into Go values */
			byt := []byte(`{"name":"Kedarnag", "emails":["abc@yopmail.com", "xyz@yopmail.com", "pqr@yopmail.com"]}`)
			var dec_data map[string]interface{}
	/* This will unmarshal the json and give us the JSON. */		
			err = json.Unmarshal(byt, &dec_data)
			fmt.Println(dec_data)

			name := dec_data["name"]
			fmt.Println("Name:", name)
	/* This will fetch me all the emails in the key emails. */
			emails := dec_data["emails"]
			fmt.Println("Emails:", emails)

	/* What if we want to get individual email of the array. As it is of interface type, we need to type cast it and convert it into Array interface*/
			email := dec_data["emails"].([]interface{})
			fmt.Println("Email:", email[0])
}



