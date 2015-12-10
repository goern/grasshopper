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

	"github.com/stretchr/testify/assert"
)

func TestDumpNuleculeRecursively(t *testing.T) {
	assert := assert.New(t)

	containerApplication, err := ParseFile("../test-fixtures/with-inherits", "yaml")

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	assert.NotNil(containerApplication)

	assert.NotEmpty(containerApplication.Dump(true))
}
