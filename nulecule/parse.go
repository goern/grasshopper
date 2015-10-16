// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package nulecule

import (
	"io"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Parse parses the Nulecule file from the given io.Reader.
func Parse(r io.Reader) (*ContainerApplication, error) {
	data, err := ioutil.ReadAll(r)

	//	log.Print(string(data))

	if err != nil {
		log.Fatal(err)
	}

	var app ContainerApplication

	unmarschalError := yaml.Unmarshal(data, app)

	if unmarschalError != nil {
		log.Fatal(unmarschalError)
	}

	return &app, unmarschalError
}

// ParseFile parses a Nulecule file at the given path.
func ParseFile(path string) (*ContainerApplication, error) {
	return nil, nil
}
