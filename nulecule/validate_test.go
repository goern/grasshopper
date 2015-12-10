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
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	assert := assert.New(t)

	containerApplication, parseError := ParseFile("../test-fixtures/Nulecule", "yaml")
	assert.Nil(parseError)

	if assert.NotNil(containerApplication) {
		assert.Equal(NuleculeVersion, containerApplication.Specversion, "Nulecule Spec Version should be 0.0.2")

		valErr := containerApplication.Validate()

		if valErr != nil {
			t.Log(valErr)
		}
	}

}

func TestInvalideNulecule(t *testing.T) {
	assert := assert.New(t)

	containerApplication, parseError := ParseFile("../test-fixtures/Nulecule", "yaml")
	assert.Nil(parseError)

	if parseError != nil {
		fmt.Printf("Error: %s", parseError.Error())
	}

	assert.NotNil(containerApplication)

	containerApplication.Specversion = "1.2.3"
	err := containerApplication.Validate()
	assert.NotNil(err)
	assert.Equal(err.(*multierror.Error).Errors[0], fmt.Errorf("'specversion' MUST be 0.0.2"))
}
