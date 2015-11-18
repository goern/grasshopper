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
	"strings"

	jww "github.com/spf13/jwalterweatherman"

	"github.com/docker/docker/builder/dockerfile/parser"
)

//GuessFromDockerfile will guess some information from a Dockerfile file
//guessing means to get all the `io.projectatomic.nulecule` labels and return
//them as a map
func GuessFromDockerfile(filename string) (*Guess, error) {
	dockerfileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		jww.ERROR.Println("failed to read the Dockerfile")
	}

	// lets parse the Dockerfile
	ast, err := parser.Parse(bytes.NewReader(dockerfileContent))

	if err != nil {
		jww.FATAL.Println("Dockerfile parse error")
	}

	//	fmt.Print(ast.Dump())
	for _, s := range guessFromLabels(ast) {
		fmt.Printf("k: %s;\t v: %s\n", s.Key, s.Value)
	}

	guesses := new(Guess)
	return guesses, nil
}

//Guess does contain all the things we guessed from a Dockerfile
type Guess struct {
	Labels []label
}

type label struct {
	Key   string
	Value string
}

func guessFromLabels(node *parser.Node) []label {
	var arr []label

	isLabel := (strings.ToUpper(node.Value) == "LABEL")

	for _, n := range node.Children {
		arr = append(arr, guessFromLabels(n)...)
	}

	var k, v string
	step := true // tick, tock
	if node.Next != nil {
		for n := node.Next; n != nil; n = n.Next {
			if len(n.Children) > 0 {
				arr = append(arr, guessFromLabels(n)...)
			} else if isLabel {
				if step {
					step = false
					// TODO: think about if you want the Quote parts
					//k = strconv.Quote(n.Value)
					k = n.Value
				} else {
					step = true
					//v = strconv.Quote(n.Value)
					v = n.Value
					arr = append(arr, label{k, v})
				}
			}
		}
	}
	return arr
}
