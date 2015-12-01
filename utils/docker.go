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

//Package utils includes all the we need for the grasshopper
package utils

import "github.com/fsouza/go-dockerclient"

//PullDockerImage will pull a Docker image and
func PullDockerImage(name string) bool {
	client, _ := docker.NewClientFromEnv()

	client.PullImage(docker.PullImageOptions{"docker.io", "goern", "grasshopper", nil, false},
		docker.AuthConfiguration{})

	return true
}
