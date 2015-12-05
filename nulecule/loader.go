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
	"archive/tar"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/fsouza/go-dockerclient"
	"github.com/hashicorp/go-multierror"
	"github.com/satori/go.uuid"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

//DockerEndpoint will hold the runtime configuration
type DockerEndpoint struct {
	Schema string // unix or tcp
	Host   string // empty if unix or hostname or ip address if tcp
	Port   int    // 0 if unix or port number if tcp
	Path   string // path to socket if unix or empty
}

var dockerEndpoint = DockerEndpoint{"unix", "", 0, "/var/run/docker.sock"}

//LoadNulecule will load a Nulecule from a URL and follow all references
// to 'external' graph components aka other Nulecules.
// It will return a fully populated ContainerApplication struct
func LoadNulecule(url *url.URL) (*ContainerApplication, error) {
	var errors *multierror.Error

	if url.Scheme != "docker" {
		jww.ERROR.Print("Grasshopper can only load Nulecules from docker:// URLs")
		errors = multierror.Append(errors, fmt.Errorf("Not a docker URL schema"))

		return nil, errors.ErrorOrNil()
	}

	// load the Nulecule from the URL
	app, err := getNuleculeFileFromDockerImage(url)
	if err != nil {
		errors = multierror.Append(errors, err)

		return nil, errors.ErrorOrNil()
	}

	// validate the Nulecule
	err = app.Validate()
	if err != nil {
		errors = multierror.Append(errors, err)

		return nil, errors.ErrorOrNil()
	}

	// figure out which graph components have a source attribute,
	// follow them and LoadNulecule()
	for _, component := range app.Graph {
		if component.Source != "" {
			jww.DEBUG.Printf("Graph Component %#v is an external Nulecule\n", component)

			componentURL, err := url.Parse(component.Source)
			// FIXME(goern) chk the err

			externalApp, err := getNuleculeFileFromDockerImage(componentURL)
			if err != nil {
				errors = multierror.Append(errors, err)

				break
			}

			// validate the external Nulecule
			err = externalApp.Validate()
			if err != nil {
				errors = multierror.Append(errors, err)

				break
			}

			jww.DEBUG.Printf("Graph Component Nulecule: %#v\n", externalApp)

			// merge loaded reference into parent ContainerApplication
		}
	}

	return app, errors.ErrorOrNil()
}

//getNuleculeFileFromDockerImage will extract and return an unvalidated
// ContainerApplication (Nulecule file)
func getNuleculeFileFromDockerImage(url *url.URL) (*ContainerApplication, error) {
	var pullImageOutputStream bytes.Buffer
	var nuleculeOutputStream bytes.Buffer
	var dockerImageName string
	var dockerRegistry = viper.GetString("dockerRegistry")

	client, err := newClientFromEndpoint(&dockerEndpoint)

	if err != nil {
		jww.ERROR.Printf("%#v\n", err)
		return nil, err
	}

	if url.Scheme == "docker" {
		dockerImageName = url.Host + url.Path
	}

	jww.DEBUG.Printf("URL path is %s\n", dockerImageName)

	// lets get the image
	err = client.PullImage(
		docker.PullImageOptions{dockerImageName, dockerRegistry, "latest", &pullImageOutputStream, false},
		docker.AuthConfiguration{})

	jww.DEBUG.Printf("pullImageOutputStream: %#v\n", pullImageOutputStream.String())

	if err != nil {
		fmt.Printf("%#v\n", err)
		return nil, err
	}

	// run that image so we can copy files from it
	containerConfig := docker.Config{
		Image: dockerImageName,
	}
	container, err := client.CreateContainer(docker.CreateContainerOptions{Name: uuid.NewV4().String(), Config: &containerConfig})

	if err != nil {
		jww.ERROR.Printf("%#v\n", err)
		return nil, err
	}
	defer client.RemoveContainer(docker.RemoveContainerOptions{ID: container.ID})

	jww.DEBUG.Printf("started %s as %s named %s\n", dockerImageName, container.ID, container.Name)

	// and copy the files (as a tar)
	err = client.CopyFromContainer(docker.CopyFromContainerOptions{&nuleculeOutputStream, container.ID, "/application-entity/Nulecule"})
	if err != nil {
		jww.ERROR.Printf("%#v\n", err)
		return nil, err
	}
	nuleculeFile := new(bytes.Buffer)
	r := bytes.NewReader(nuleculeOutputStream.Bytes())

	// what we downloaded is a tar, so lets get the real file!
	tr := tar.NewReader(r)
	tr.Next()
	if err != nil && err != io.EOF {
		jww.ERROR.Printf("the tar archive copied form the container seems to be broken: %s", err)
	}

	if _, err := io.Copy(nuleculeFile, tr); err != nil {
		jww.ERROR.Printf("Can't copy the Nulecule file: %s", err)
	}

	jww.DEBUG.Printf("get Nuleculde from %s:\n%s\n", container.ID, nuleculeFile)

	app, err := Parse(nuleculeFile)

	return app, err
}

//getArtifactsFromDockerImage will get all artifact files for all providers
// and return them as a tar archive
func getArtifactsFromDockerImage(URL string) (*bytes.Buffer, error) {
	splitURL := strings.Split(URL, "/")
	var artifactsOutputStream bytes.Buffer
	var errors *multierror.Error

	client, err := newClientFromEndpoint(&dockerEndpoint)

	if err != nil {
		jww.ERROR.Printf("%#v\n", err)
		errors = multierror.Append(errors, err)
		return nil, errors.ErrorOrNil()
	}

	// run that image so we can copy files from it
	containerConfig := docker.Config{
		Image: splitURL[2] + "/" + splitURL[3],
	}
	container, err := client.CreateContainer(docker.CreateContainerOptions{Name: uuid.NewV4().String(), Config: &containerConfig})
	defer client.RemoveContainer(docker.RemoveContainerOptions{ID: container.ID})

	if err != nil {
		jww.ERROR.Printf("%#v\n", err)
		errors = multierror.Append(errors, err)
		return nil, errors.ErrorOrNil()
	}

	// and copy the files (as a tar)
	err = client.CopyFromContainer(docker.CopyFromContainerOptions{&artifactsOutputStream, container.ID, "/application-entity/Nulecule"})
	if err != nil {
		jww.ERROR.Printf("%#v\n", err)
		errors = multierror.Append(errors, err)
		return nil, errors.ErrorOrNil()
	}

	return &artifactsOutputStream, errors.ErrorOrNil()
}

func newClientFromEndpoint(endpoint *DockerEndpoint) (*docker.Client, error) {
	var client *docker.Client
	var err error

	if endpoint == nil {
		endpoint = &dockerEndpoint
	}

	if (os.Getenv("DOCKER_HOST") != "") &&
		(os.Getenv("DOCKER_TLS_VERIFY") == "1") &&
		(os.Getenv("DOCKER_CERT_PATH") != "") { // FIXME this seems to be a hack
		client, err = docker.NewClientFromEnv()
	} else if endpoint.Schema == "unix" {
		client, err = docker.NewClient(endpoint.Schema + "://" + endpoint.Path)
	} else {
		return nil, errors.New("invalid docker endpoint")
	}

	return client, err
}
