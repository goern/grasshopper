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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	assert := assert.New(t)

	containerApplication, parseError := ParseFile("../test-fixtures/Nulecule")

	if parseError != nil {
		t.Log(parseError)
		// FIXME this should panic!
	}

	assert.NotNil(containerApplication)
}

func TestParseFileWithInherits(t *testing.T) {
	assert := assert.New(t)

	containerApplication, parseError := ParseFile("../test-fixtures/with-inherits")

	if parseError != nil {
		t.Log(parseError)
		// FIXME this should panic!

	}

	assert.NotNil(containerApplication)
}

func TestParseFileWithConstraints(t *testing.T) {
	assert := assert.New(t)

	containerApplication, parseError := ParseFile("../test-fixtures/with-constraints")

	if parseError != nil {
		t.Log(parseError)
		// FIXME this should panic!

	}

	assert.NotNil(containerApplication)

	//	t.Log(string(containerApplication))
}
