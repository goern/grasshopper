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
	"net/http"

	jww "github.com/spf13/jwalterweatherman"
)

//LoadNulecule will load a Nulecule from a URL and follow all references
// to 'external' graph components aka other Nulecules.
// It will return a fully populated ContainerApplication struct
func LoadNulecule(URL string) (*ContainerApplication, error) {
	// load the Nulecule from the URL
	resp, err := http.Get(URL)
	if err != nil {
		jww.FATAL.Printf("cant load Nulecule file from %s", URL)
	}

	defer resp.Body.Close()
	app, err := Parse(resp.Body)

	// figure out which graph components have a source attribute,
	// follow them and LoadNulecule()

	// merge loaded reference into parent ContainerApplication

	return app, err
}
