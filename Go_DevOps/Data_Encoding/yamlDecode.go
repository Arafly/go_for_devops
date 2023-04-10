package main

import (
	"github.com/go-yaml/yaml"
	"fmt"
	"os"
	"log"
)

var yamlFile = []byte(`
	User: John Doak
	ID: 25
T	raits: ["Tall", "Blonde", "Dashing"]
`)

func main() {
	// r := strings.NewReader(`
	// 	- name: John Raymond
	// 		surname: Legrasse
	// 		job: policeman
	// 	- name: "Francis"
	// 		surname: Wayland Thurston
	// 	job: anthropologist`)
	// // define a new decoder
	// d := yaml.NewDecoder(r)
	// var c []Character
	// // decode the reader
	// if err := d.Decode(&c); err != nil {
	// 	log.Fatalln(err)
	// } log.Printf("%+v", c)


	// Unmarshalling into map
	data := map[string]interface{}{}

	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		panic(err)
	}

	fmt.Println(data["User"])
	fmt.Println(data["ID"])

	// We check that this key is set so that our type assertion below will not panic.
	if _, ok := data["Traits"]; ok {
		// Because YAML can store lists of different types, all arrays are []interface{} when
		// decoding into a map.
		for _, trait := range data["Traits"].([]interface{}) {
			fmt.Println("Trait: ", trait)
		}
	}

}