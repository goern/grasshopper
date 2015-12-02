/*
 Copyright 2015 Red Hat, Inc.

 This file is part of Grasshopper.

 Grasshopper is free software: you can redistribute it and/or modify
 it under the terms of the GNU Lesser General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 Grasshopper is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Lesser General Public License for more details.

 You should have received a copy of the GNU Lesser General Public License
 along with Grasshopper. If not, see <http://www.gnu.org/licenses/>.
*/

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
// the Nulecule app with the data read in.
func FollowReference(URL string, app *ContainerApplication) error {
	jww.DEBUG.Printf("following a reference to %s\n", URL)

	return nil
}
