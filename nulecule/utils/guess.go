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

// Package utils provides some useful utility functions for Dockerfiles or
// Nulecule files.
package utils

import (
	"bytes"
	"fmt"
	"go/ast"
	"io/ioutil"

	jww "github.com/spf13/jwalterweatherman"

	"github.com/docker/docker/builder/dockerfile/parser"
)

//GuessFromDockerfile will guess some information from a Dockerfile file
//guessing means to get all the `io.projectatomic.nulecule` labels and return
//them as a map
func GuessFromDockerfile(filename string) (map[string]string, error) {
	labels := make(map[string]string)

	dockerfileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		jww.ERROR.Println("failed to read the Dockerfile")
	}

	// lets parse the Dockerfile
	root, err := parser.Parse(bytes.NewReader(dockerfileContent))

	fmt.Print(string(root.Dump()))

	// ast.Walk(VisitorFunc(labelVisitor), root)

	return labels, nil
}

type visitorFunc func(n ast.Node) ast.Visitor

func labelVisitor(n ast.Node) (w ast.Visitor) {
	return nil
}
