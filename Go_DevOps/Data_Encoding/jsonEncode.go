// We can encode JSON with a json.Marshal function that takes a byte slice
// We could also use the json.Encoder type that uses io.Writer instead.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	Id int `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


func main() {
	//  struct to JSON (using json.Marshal)
	emp := Employee{
		Id: 135,
		Name: "John", 
		Email: "john.doe@gmail.com",
	}
	b, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Stdout.Write(b)

	// using json.Encoder
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "\t")
	c := Employee{
		Name: "Charles Dexter",
		Email: "jackson@mail.com",
	} if err := e.Encode(c); err != nil {
		log.Fatalln(err)
	}

	// Marshalling a map into JSON
	// We then create a map of string keys and interface{} values. This is a common pattern when you need to handle multiple message types or do discovery on a message.

	// data := map[string]interface{}{}
	// if err := json.Marshal(data); err != nil {
	// 	return err
	// }
}
