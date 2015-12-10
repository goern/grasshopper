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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/go-multierror"
)

func TestLoadNuleculeWithNonDockerURL(t *testing.T) {
	assert := assert.New(t)

	url, _ := url.Parse("http://example.com/Nulecule")

	app, err := LoadNulecule(&DefaultLoaderOptions, url)

	assert.NotNil(err)

	assert.Equal(err.(*multierror.Error).Errors[0], fmt.Errorf("Not a docker URL schema"))
	assert.Nil(app)
}

func testGetNuleculeFileFromDockerImage(t *testing.T) {
	assert := assert.New(t)

	// This should work, thus no error checking
	url, _ := url.Parse("docker://projectatomic/wordpress-centos7-atomicapp")

	app, err := getNuleculeFileFromDockerImage(&DefaultLoaderOptions, url)

	assert.Nil(err)
	assert.NotNil(app)
}

func testGetArtifactsFromDockerImage(t *testing.T) {
	assert := assert.New(t)

	// This should work too, thus no error checking ;)
	url, _ := url.Parse("docker://projectatomic/wordpress-centos7-atomicapp")

	app, err := getArtifactsFromDockerImage(&DefaultLoaderOptions, url)

	assert.Nil(err)
	assert.NotNil(app)
}
