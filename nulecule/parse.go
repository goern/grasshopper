// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package nulecule

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Parse parses the Nulecule file from the given io.Reader.
func Parse(r io.Reader) (*ContainerApplication, error) {
	data, err := ioutil.ReadAll(r)

	if err != nil {
		log.Fatal(err)
	}

	app := ContainerApplication{}

	unmarschalError := yaml.Unmarshal(data, &app)

	if unmarschalError != nil {
		log.Fatal(unmarschalError)
	}

	// TODO before returning we should do some sanity checks, like: specversion equals grasshopper supported spec

	return &app, unmarschalError
}

// ParseFile parses a Nulecule file at the given path.
func ParseFile(filename string) (*ContainerApplication, error) {
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	return Parse(f)
}
