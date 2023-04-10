// Just like JSON with a Marshal function that takes a byte slice or an Encoder type that uses io.Writer instead.

package main

import (
	"log"
	"os"
	"github.com/go-yaml/yaml"
)

type Character struct {
	Name string `yaml:"name"`
	Surname string `yaml:"surname"`
	Job string `yaml:"job,omitempty"`
	YearOfBirth int `yaml:"year_of_birth,omitempty"`
}

func main() {
	// We can use them to encode a data structure or, in this case, a list of structures:
	var chars = []Character{{
		Name: "William",
		Surname: "Dyer",
		Job: "professor",
		YearOfBirth: 1875,
	}, {
		Surname: "Danforth",
		Job: "student",
	}}
	e := yaml.NewEncoder(os.Stdout)
	if err := e.Encode(chars); err != nil {
		log.Fatalln(err)
	}

	// Marshalling into map
	// if err := yaml.Unmarshal(data); err != nil {
	// 	log.Fatalln(err)
	// }
}
