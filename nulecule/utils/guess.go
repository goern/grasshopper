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
	"log"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"

	"github.com/docker/docker/builder/dockerfile/parser"
)

//NuleculePersistentVolumeTemplate is a template to generate a YAML http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#storageRequirementsObject snippet
const NuleculePersistentVolumeTemplate = `---
- persistentVolume:
    name: "{{.Name}}"
    accessMode: "ReadWrite"
{{if gt .Size 0}}    size: {{.Size}}{{else}}    size: 4 # GB by default{{end}}
`

//NuleculePersistentVolume is the specification for a http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#storageRequirementsObject
type NuleculePersistentVolume struct {
	Name string
	Size int
}

//GuessFromDockerfile will guess some information from a Dockerfile file
//guessing means to get all the `io.projectatomic.nulecule` labels and return
//them as a map
func GuessFromDockerfile(filename string) (Guess, error) {
	dockerfileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		jww.ERROR.Println("failed to read the Dockerfile")
	}

	// lets parse the Dockerfile
	ast, err := parser.Parse(bytes.NewReader(dockerfileContent))

	if err != nil {
		jww.FATAL.Println("Dockerfile parse error")
	}

	for _, s := range guessFromLabels(ast) {
		jww.DEBUG.Printf("k: %s;\t v: %s\n", s.Key, s.Value)
	}

	if viper.GetBool("Experimental") {
		snippetsFromLabels(ast)
	}

	var guesses Guess
	guesses.Labels = guessFromLabels(ast)
	return guesses, err
}

//Guess does contain all the things we guessed from a Dockerfile
type Guess struct {
	Labels []Label
}

//Label is a generic structure holding LABELs (of Dockerfiles)
type Label struct {
	Key   string
	Value string
}

func guessFromLabels(node *parser.Node) []Label {
	var arr []Label

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
					k = strconv.Quote(n.Value)
				} else {
					step = true
					v = n.Value
					arr = append(arr, Label{k, v})
				}
			}
		}
	}
	return arr
}

//SpaceMap will use strings.Map() to map spaces to nothing
// see https://stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string-in-golang?answertab=votes#tab-top
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func stringToArrayOfStrings(str string) []string {
	return strings.Split(SpaceMap(strings.Replace(str, "\"", "", -1)), ",")
}

func generateNuleculePersistentVolume(spec NuleculePersistentVolume) string {
	var buffer bytes.Buffer

	// Create a new template and parse the NuleculePersistentVolumeTemplate into it.
	t := template.Must(template.New("PersistentVolume").Parse(NuleculePersistentVolumeTemplate))

	spec.Name = strings.Replace(strings.Replace(spec.Name, "/", "-", -1), "-", "", 1) // there shall be no / in a name

	err := t.Execute(&buffer, spec)
	if err != nil {
		log.Println("executing template:", err)
	}

	return buffer.String()
}

func snippetsFromLabels(node *parser.Node) {
	isLabel := (strings.ToUpper(node.Value) == "LABEL")

	for _, n := range node.Children {
		snippetsFromLabels(n)
	}

	if node.Next != nil {
		for n := node.Next; n != nil; n = n.Next {
			if len(n.Children) > 0 {
				snippetsFromLabels(n)
			} else if isLabel {
				switch strings.ToLower(n.Value) {
				case "io.k8s.description":
					fmt.Printf("this is a io.k8s.description LABEL, content: %s\n", n.Next.Value)
				case "io.k8s.display-name":
					fmt.Printf("this is a io.k8s.display-name LABEL, content: %s\n", n.Next.Value)
				case "io.projectatomic.nulecule.volume.data":
					fmt.Printf("found a io.projectatomic.nulecule.volume.data LABEL and will genrate a storage claim\n%s\n",
						generateNuleculePersistentVolume(NuleculePersistentVolume{strings.Replace(n.Next.Value, "\"", "", -1), 1})) // TODO doc that this is GiB
				case "io.projectatomic.nulecule.environment.required":
					fmt.Printf("this is a io.projectatomic.nulecule.environment.required LABEL, content: %#v\n", stringToArrayOfStrings(n.Next.Value))
				case "io.projectatomic.nulecule.environment.optional":
					fmt.Printf("this is a io.projectatomic.nulecule.environment.optional LABEL, content: %#v\n", stringToArrayOfStrings(n.Next.Value))
				}

			}
		}
	}
}
