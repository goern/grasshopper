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
	"fmt"
	"net/url"

	"github.com/hashicorp/go-multierror"
	"github.com/xeipuuv/gojsonschema"

	jww "github.com/spf13/jwalterweatherman"
)

// SpecVersion specifies the Nulecule Specification's version to validate with
var SpecVersion string

var schemaLocation = map[string]string{
	"0.0.2": "http://goern.github.io/grasshopper/nulecule/spec/0.0.2/schema.json",
}

//ValidateSchema will validate a file with a Nulecule Specification
func ValidateSchema(schemaVersion string, location *url.URL) (bool, error) {
	var rc error

	// check if schemaVersion equals nulecule.NuleculeReleasedVersions
	if schemaVersion != NuleculeReleasedVersions {
		return false, fmt.Errorf("The specified version (%s) of the Nulecule Specification is invalid", schemaVersion)
	}

	schemaLoader := gojsonschema.NewReferenceLoader(schemaLocation[schemaVersion])
	documentLoader := gojsonschema.NewReferenceLoader(location.String())

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false, err
	}

	if result.Valid() {
		// until now, all is good
		// so lets have a look at the artifacts

		return true, nil
	}

	fmt.Printf("The document is not valid. see errors :\n")
	for _, desc := range result.Errors() {
		rc = multierror.Append(rc, fmt.Errorf("%s\n", desc.Description()))
		fmt.Printf("- %s\n", desc)
	}

	rc = multierror.Append(fmt.Errorf("The document is not valid with Nulecule Specification version %s", schemaVersion))

	return false, rc
}

// Validate validates the Nulecule file
func (nulecule *ContainerApplication) Validate() error {
	var result *multierror.Error

	// lets see if we are supposed to work on a 0.0.2 Nulecule
	if nulecule.Specversion != "0.0.2" {
		result = multierror.Append(result, fmt.Errorf(
			"'specversion' MUST be 0.0.2"))
	}

	for _, component := range nulecule.Graph {
		for provider, artifact := range component.Artifacts {
			jww.DEBUG.Printf("Component %s, looking for Artifacts of %s Provider: %#v\n", component.Name, provider, artifact)
		}

	}

	return result.ErrorOrNil()
}
