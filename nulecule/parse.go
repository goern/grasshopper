// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package nulecule

import (
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	jww "github.com/spf13/jwalterweatherman"
)

// Parse parses the Nulecule file from the given io.Reader.
func Parse(r io.Reader) (*ContainerApplication, error) {
	data, err := ioutil.ReadAll(r)

	if err != nil {
		jww.FATAL.Println(err)
		return nil, err
	}

	app := ContainerApplication{}

	unmarschalError := yaml.Unmarshal(data, &app)

	if unmarschalError != nil {
		jww.ERROR.Println(unmarschalError)
		return nil, unmarschalError // FIXME ERROR: 2015/11/21 yaml: unmarshal errors: line 18: cannot unmarshal !!map into string

	}

	// TODO before returning we should do some sanity checks, like: specversion equals grasshopper supported spec

	return &app, unmarschalError
}

// ParseFile parses a Nulecule file at the given path.
func ParseFile(filename string) (*ContainerApplication, error) {
	f, err := os.Open(filename)

	if err != nil {
		jww.FATAL.Println(err)
		return nil, err
	}

	return Parse(f)
}

// FollowReference will follow a reference to external Nulecules and populate
// the Nulecule nulecule with the data read in.
func FollowReference(filename string, nulecule *ContainerApplication) error {
	// TODO
	return nil
}
