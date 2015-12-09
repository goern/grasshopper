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

package utils

import (
	"fmt"

	"github.com/goern/grasshopper/nulecule"
	"github.com/xeipuuv/gojsonschema"
)

// SpecVersion specifies the Nulecule Specification's version to validate with
var SpecVersion string

var schemaLocation = map[string]string{
	"0.0.2": "http://goern.github.io/grasshopper/nulecule/spec/0.0.2/schema.json",
}

//ValidateFile will validate a file with a Nulecule Specification
func ValidateFile(schemaVersion, something string) (bool, error) {
	// check if schemaVersion equals nulecule.NuleculeReleasedVersions
	if schemaVersion != nulecule.NuleculeReleasedVersions {
		return false, fmt.Errorf("The specified version (%s) of the Nulecule Specification is invalid", schemaVersion)
	}

	schemaLoader := gojsonschema.NewReferenceLoader(schemaLocation[schemaVersion])
	documentLoader := gojsonschema.NewReferenceLoader(something)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false, err
	}

	if result.Valid() {
		return true, nil
	}

	fmt.Printf("The document is not valid. see errors :\n")
	for _, desc := range result.Errors() {
		fmt.Printf("- %s\n", desc)
	}

	return false, fmt.Errorf("The document is not valid with Nulecule Specification version %s", schemaVersion)
}
