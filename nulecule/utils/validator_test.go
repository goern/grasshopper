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
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	assert := assert.New(t)

	location, err := url.Parse("http://goern.github.io/grasshopper/nulecule/spec/0.0.2/a-fixture-Nulecule")
	assert.Nil(err)

	valid, err := Validate("0.0.2", location)
	assert.Nil(err)
	assert.True(valid)

}

func TestValidateWithUnknownSpecificationVersion(t *testing.T) {
	assert := assert.New(t)

	location, err := url.Parse("http://goern.github.io/grasshopper/nulecule/spec/0.0.2/a-fixture-Nulecule")
	assert.Nil(err)

	valid, err := Validate("0.0.99", location)
	assert.NotNil(err)
	assert.Equal("The specified version (0.0.99) of the Nulecule Specification is invalid", err.Error())
	assert.False(valid)

}

func TestValidateWithYAMLNulecule(t *testing.T) {
	assert := assert.New(t)

	location, err := url.Parse("https://raw.githubusercontent.com/projectatomic/nulecule-library/master/flask-redis-centos7-atomicapp/Nulecule")
	assert.Nil(err)

	valid, err := Validate("0.0.2", location)
	assert.NotNil(err)
	fmt.Printf("%#v\n", err)
	assert.False(valid)

}
