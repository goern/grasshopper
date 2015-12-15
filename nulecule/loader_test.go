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

func testGetNuleculeFromDockerImage(t *testing.T) {
	assert := assert.New(t)

	// This should work, thus no error checking
	url, _ := url.Parse("docker://projectatomic/wordpress-centos7-atomicapp")

	app, err := getNuleculeFromDockerImage(&DefaultLoaderOptions, url)

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

func TestLoadNuleculeWithErrors(t *testing.T) {
	assert := assert.New(t)

	getNuleculeFromDockerImage = func(options *LoaderOptions, url *url.URL) (*ContainerApplication, error) {
		return nil, fmt.Errorf("Not a docker URL schema")
	}

	// This should work, thus no error checking
	url, _ := url.Parse("docker://projectatomic/wordpress-centos7-atomicapp")

	app, err := LoadNulecule(&DefaultLoaderOptions, url)

	assert.NotNil(err)
	assert.Nil(app)

}

func TestLoadNuleculeWithValidNulecule(t *testing.T) {
	assert := assert.New(t)

	getNuleculeFromDockerImage = func(options *LoaderOptions, url *url.URL) (*ContainerApplication, error) {
		return ParseFile("../test-fixtures/Nulecule")
	}

	// This should work, thus no error checking
	url, _ := url.Parse("docker://projectatomic/wordpress-centos7-atomicapp")

	app, err := LoadNulecule(&DefaultLoaderOptions, url)

	assert.Nil(err)
	assert.NotNil(app)
}

func TestLoadNuleculeWithInalidNulecule(t *testing.T) {
	assert := assert.New(t)

	getNuleculeFromDockerImage = func(options *LoaderOptions, url *url.URL) (*ContainerApplication, error) {
		app, err := ParseFile("../test-fixtures/Nulecule")
		app.Specversion = "0.0.99" // invalidate the ContainerApplication

		return app, err
	}

	// This should work, thus no error checking
	url, _ := url.Parse("docker://projectatomic/wordpress-centos7-atomicapp")

	app, err := LoadNulecule(&DefaultLoaderOptions, url)

	assert.NotNil(err)
	assert.Nil(app)
}
