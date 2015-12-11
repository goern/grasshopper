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

package nulecule

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// ParseError denotes failing to parse Nulecule file.
type ParseError struct {
	err error
}

// Returns the formatted Nulecule file parser error.
func (pe ParseError) Error() string {
	return fmt.Sprintf("While parsing Nulecule file: %s", pe.err.Error())
}

// Parse parses the Nulecule file from the given io.Reader.
func Parse(r io.Reader) (*ContainerApplication, error) {
	data, err := ioutil.ReadAll(r)

	format := guessFileFormat(r)

	if err != nil {
		return nil, ParseError{err}
	}

	app := ContainerApplication{}

	switch strings.ToLower(format) {
	case "yaml", "yml", "application/x-yaml", "text/x-yaml; charset=utf-8", "application/octet-stream":
		unmarschalError := yaml.Unmarshal(data, &app)

		if unmarschalError != nil {
			fmt.Printf("Parser: %s\n", unmarschalError.Error())
			return nil, ParseError{unmarschalError}
		}
	case "json", "application/javascript", "text/plain; charset=utf-8":
		unmarschalError := json.Unmarshal(data, &app)

		if unmarschalError != nil {
			fmt.Printf("Parser: %s\n", unmarschalError.Error())
			return nil, ParseError{unmarschalError}
		}
	default:
		return nil, ParseError{fmt.Errorf("File format %s is not supported", format)}
	}

	return &app, nil
}

// ParseFile parses a Nulecule file at the given path.
func ParseFile(filename string) (*ContainerApplication, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, ParseError{err}
	}

	return Parse(f)
}

func guessFileFormat(r io.Reader) string {
	buf := make([]byte, 512)

	r.Read(buf)
	return http.DetectContentType(buf)
}
