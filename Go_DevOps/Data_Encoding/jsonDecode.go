// There two ways of decoding data in JSON—
// The first is the json.Unmarshal: a function that uses a byte slice as input
// The second is the json.Decoder type that uses a generic io.Reader to get the encoded content.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Character struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Job         string `json:"job,omitempty"`
	YearOfBirth int    `json:"year_of_birth,omitempty"`  // the omitempty tag that is used for output purposes to avoid marshaling the field if it has a zero value.
}



func main() {
	// JSON to struct (using json.Unmarshal)
	emp := Employee{}
	b := []byte(`{"id":135,"name":"John","email":"john@gmail"}`)
	if err := json.Unmarshal(b, &emp); err != nil {
		fmt.Println(err)
		return
	}
	
	// using json.Decoder
	r := strings.NewReader(`{"name":"Lavinia","surname":"Whateley","year_of_birth":1878}`)
	d := json.NewDecoder(r)
	var c Character
	if err := d.Decode(&c); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", c)


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