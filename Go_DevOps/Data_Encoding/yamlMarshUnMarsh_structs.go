package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Jobs []Job
}

type Job struct {
	Name     string
	Interval time.Duration
	Cmd      string
}

func main() {
	// Marshalling from structs
	c := Config{
		Jobs: []Job{
			{
				Name:     "Clear tmp",
				Interval: 24 * time.Hour,
				Cmd:      "rm -rf " + os.TempDir(),
			},
		},
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)

	// Unmarshalling into structs

	data := []byte(`
		jobs:
		- name: Clear tmp
		  interval: 24h
		  cmd: rm -rf ` + os.TempDir() + `
	`)
	c2 := Config{}
	if err := yaml.Unmarshal(data, &c2); err != nil {
		panic(err)
	}
	for _, job := range c2.Jobs {
		fmt.Printf("Name: ", job.Name)
		fmt.Println("Interval: ", job.Interval)
	}
}