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
	"errors"
	"fmt"
	"strings"

	"github.com/fsouza/go-dockerclient"

	jww "github.com/spf13/jwalterweatherman"
)

//LoadNulecule will load a Nulecule from a URL and follow all references
// to 'external' graph components aka other Nulecules.
// It will return a fully populated ContainerApplication struct
func LoadNulecule(URL string) (*ContainerApplication, error) {
	if !strings.HasPrefix(URL, "docker://") {
		jww.WARN.Print("Grasshopper can only load Nulecules from docker:// URLs")
		return nil, errors.New("Not a docker URL schema")
	}

	// load the Nulecule from the URL
	app, err := getNuleculeFileFromDockerImage(URL)

	// figure out which graph components have a source attribute,
	// follow them and LoadNulecule()

	// merge loaded reference into parent ContainerApplication

	return app, err
}

func getNuleculeFileFromDockerImage(URL string) (*ContainerApplication, error) {
	// docker://registry/repository/image:tag
	splitURL := strings.Split(URL, "/")

	client, err := docker.NewClient("unix:///var/run/docker.sock")

	if err != nil {
		jww.ERROR.Printf("%#v\n", err)
		return nil, err
	}

	// TODO make registry a config option
	err = client.PullImage(
		docker.PullImageOptions{splitURL[2] + "/" + splitURL[3], "registry.docker.com", "latest", nil, false},
		docker.AuthConfiguration{})

	if err != nil {
		fmt.Printf("%#v\n", err)
		return nil, err
	}

	return nil, err
}
