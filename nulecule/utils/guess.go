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
	"io/ioutil"
	"strconv"
	"strings"

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
	ast, err := parser.Parse(bytes.NewReader(dockerfileContent))

	if err != nil {
		jww.FATAL.Println("Dockerfile parse error")
	}

	fmt.Print(ast.Dump())

	cursor := ast
	var n int
	for cursor.Next != nil {
		cursor = cursor.Next
		jww.DEBUG.Printf("cursor.Value = %s\n", cursor.Value)
		n++
	}
	msgList := make([]string, n)
	strList := []string{}
	msg := ""
	jww.DEBUG.Printf("size of msgList is %d\n", n)

	var i int
	for ast.Next != nil {
		ast = ast.Next
		var str string
		str = ast.Value

		strList = append(strList, str)
		msgList[i] = ast.Value
		i++
	}

	msg += " " + strings.Join(msgList, " ")
	fmt.Println(msg)

	return labels, nil
}

type Guesses struct {
	Labels map[string][]string
}

func printLabels(node *parser.Node) string {
	str := ""

	if node.Value == "label" {
		if node.Next != nil {
			for n := node.Next; n != nil; n = n.Next {
				if len(n.Children) > 0 {
					str += printLabels(n)
				} else {
					str += " " + strconv.Quote(n.Value)
				}
			}
		}
	}

	for _, n := range node.Children {
		str += printLabels(n) + "\n"
	}

	return strings.TrimSpace(str)
}
