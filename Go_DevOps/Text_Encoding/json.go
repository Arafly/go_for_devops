// Marshal and Unmarshal JSON
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	Id int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


// func main() {
// 	//  struct to JSON
// 	emp := Employee{
	// 	Id: 135,
	// 	Name: "John", 
	// 	Email: "john.doe@gmail.com"
	// }
// 	b, err := json.Marshal(emp)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	os.Stdout.Write(b)

// 	// Marshalling a map into JSON
// 	// We then create a map of string keys and interface{} values. This is a common pattern when you need to handle multiple message types or do discovery on a message.

// 	data := map[string]interface{}{}
// 	if err := json.Marshal(data); err != nil {
// 		return err
// 	}
// }

func main() {
	// JSON to struct
	emp := Employee{}
	b := []byte(`{"id":135,"name":"John","email":"john@gmail"}`)
	if err := json.Unmarshal(b, &emp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(emp)
	


	// When you need to handle multiple message types or do discovery on a message, Go allows you to decode messages into map[string]interface{}, where the string key represents the field name and interface{} represents the value.
	
	// Let's examine an example of unmarshaling a file into a map:

	// Read content of the data.json andcreates a map, called data, to store our JSON content.
	// Unmarshals the raw bytes representing the JSON into data.
	b, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := map[string]interface{}{}
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Println(err)
		return
	}
	// It looks up the user key in data. If user does not exist, we return an error.
	// If it does exist, we type assert to determine what the value type is.
	// If the value is a string, return the content. If otherwis,return an error.
	v, ok := data["user"]
	if !ok {
    	return "", errors.New("json does not contain key 'user'")
	}
	switch user := v.(type) {
	case string:
     	return user, nil
	}
	return "", fmt.Errorf("key 'user' is not a string, was %T", v)
}