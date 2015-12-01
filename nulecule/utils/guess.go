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
	"text/template"
	"unicode"

	"github.com/goern/grasshopper/nulecule"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"

	"github.com/docker/docker/builder/dockerfile/parser"
)

//NuleculeParameterTemplate is a template to generate a YAML http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#paramsObject snippet
const NuleculeParameterTemplate = `      - name: "{{.Name}}"
{{if .Description}}        description: "{{.Description}}"{{else}}{{if .Required}}        description: "{{.Name}} is a REQUIRED parameter"{{end}}{{end}}
{{if .Default}}        default: "{{.Default}}"{{end}}
`

// FIXME indent level will be a problem here

//NuleculePersistentVolumeTemplate is a template to generate a YAML http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#storageRequirementsObject snippet
const NuleculePersistentVolumeTemplate = `  - persistentVolume:
      name: "{{.Name}}"
      accessMode: "ReadWrite"
{{if .Size }}      size: "{{.Size}}"{{else}}      size: "4Gi" # GB by default{{end}}
`

//NuleculePersistentVolume is the specification for a http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#storageRequirementsObject
type NuleculePersistentVolume struct { // FIXME cant this be StorageRequirement from base?
	Name string
	//  	AccessMode string
	Size string // FIXME int?? long?
}

//NuleculeMetadataTemplate is a template to generate a YAML http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#metadataObject snippet
const NuleculeMetadataTemplate = `metadata:
  name: "{{.Name}}"
{{ if .Version }}  appversion: "{{.Version}}"{{ end }}
{{ if .Description }}  description: "{{.Description}}"{{ end }}
` // FIXME add a license object to the template

//NuleculeMetadata is the specification for a http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#metadataObject
type NuleculeMetadata struct {
	Name, Version, Description, License string
}

//GuessFromDockerfile will guess some information from a Dockerfile file
//guessing means to get all the LABELs and process them somehow
func GuessFromDockerfile(filename string) (map[string]string, string, error) {
	dockerfileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		jww.ERROR.Println("failed to read the Dockerfile")
		return nil, "", err
	}

	// lets parse the Dockerfile
	ast, err := parser.Parse(bytes.NewReader(dockerfileContent))

	if err != nil {
		jww.FATAL.Println("Dockerfile parse error")
		return nil, "", err
	}

	var result map[string]string
	result = make(map[string]string)

	for _, s := range guessFromLabels(ast) {
		result[strings.Replace(s.Key, "\"", "", -1)] = strings.Replace(s.Value, "\"", "", -1)
	}

	for key, value := range result {
		jww.DEBUG.Printf("result: LABEL: %s;\t VALUE: %s\n", key, value)
	}

	// if --experimental is true, we print all snippets to STDOUT
	resultingNulecule := ""
	if viper.GetBool("Experimental") {
		resultingNulecule = snippetsFromLabelsMap(result)
	}

	return result, resultingNulecule, err
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

func generateNuleculePersistentVolume(spec NuleculePersistentVolume) string {
	var buffer bytes.Buffer

	// Create a new template and parse the NuleculePersistentVolumeTemplate into it.
	t := template.Must(template.New("PersistentVolume").Parse(NuleculePersistentVolumeTemplate))

	spec.Name = strings.Replace(spec.Name, "/", "-", -1) // there shall be no / in a name

	err := t.Execute(&buffer, spec)
	if err != nil {
		jww.ERROR.Println("generating Nulecule PersistentVolume snippet:", err)
		return ""
	}

	return buffer.String()
}

func generateNuleculeMetadata(spec NuleculeMetadata) string {
	var buffer bytes.Buffer

	// Create a new template and parse the NuleculePersistentVolumeTemplate into it.
	t := template.Must(template.New("Metadata").Parse(NuleculeMetadataTemplate))

	err := t.Execute(&buffer, spec)
	if err != nil {
		jww.ERROR.Println("generating Nulecule Metadata snippet:", err)
		return ""
	}

	return buffer.String()
}

func generateNuleculeParameter(spec nulecule.Param) string {
	var buffer bytes.Buffer

	// Create a new template and parse the NuleculeParameterTemplate into it.
	// TODO set indent level based on position in graph
	t := template.Must(template.New("Parameter").Parse(NuleculeParameterTemplate))

	err := t.Execute(&buffer, spec)
	if err != nil {
		jww.ERROR.Println("generating Nulecule Parameter snippet:", err)
		return ""
	}

	return buffer.String()
}

//InLabels will test it a label is in a LABELs map (derived from Dockerfile)
func InLabels(label string, labels map[string]string) bool {
	for key := range labels {
		if strings.ToUpper(key) == strings.ToUpper(label) {
			return true
		}
	}

	return false
}

//GetNuleculeVolumesFromLabels will return a list of Nulecule volumes found
// within a Dockerfile. It will return a list containing
// NuleculePersistentVolume. All strings are converted to lower case letters
func GetNuleculeVolumesFromLabels(labels map[string]string) []NuleculePersistentVolume {
	var result []NuleculePersistentVolume

	for key, value := range labels {
		parts := strings.Split(strings.ToLower(key), ".")
		if strings.Join(parts[0:len(parts)-1], ".") == strings.ToLower("io.projectatomic.nulecule.volume") {
			jww.DEBUG.Printf("got a Volume: %s, path should be %s\n", parts[len(parts)-1], value)

			splitted := strings.Split(value, ",") // the will split path and size

			if len(splitted) < 2 {
				result = append(result, NuleculePersistentVolume{parts[len(parts)-1], "4Gi"})
			} else {
				result = append(result, NuleculePersistentVolume{parts[len(parts)-1], SpaceMap(splitted[1])})
			}
		}
	}

	return result
}

//GetNuleculeParametersFromLabels will return a list of Nulecule parameters found
// within a Dockerfile. It will return a list containing Param.
// All strings are converted to lower case letters
func GetNuleculeParametersFromLabels(labels map[string]string) []nulecule.Param {
	var result []nulecule.Param

	for key, value := range labels {
		parts := strings.Split(strings.ToLower(key), ".")
		if strings.Join(parts[0:len(parts)-1], ".") == strings.ToLower("io.projectatomic.nulecule.environment") {
			// we need to split the value by , and generate for each items a parameter
			params := strings.Split(strings.ToUpper(value), ",")

			for _, param := range params {
				jww.DEBUG.Printf("got Parameter: %s, %s\n", SpaceMap(param), parts[len(parts)-1])

				result = append(result, nulecule.Param{SpaceMap(param), "", nulecule.Constraint{}, "", isRequired(parts[len(parts)-1])})

			}
		}
	}

	return result
}

// This is the real thing, it will generate the Nulecule file from all the snippets we have gathered.
func snippetsFromLabelsMap(labels map[string]string) string {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, `---
specversion: 0.0.2
id: %s

`, generateNuleculeID(labels["io.k8s.display-name"]))

	// gathers some data to start with
	if InLabels("io.k8s.description", labels) && InLabels("io.k8s.display-name", labels) {
		versionString := ""

		if labels["Release"] != "" {
			versionString = labels["Version"] + "-" + labels["Release"]
		} else {
			versionString = labels["Version"]
		}

		jww.DEBUG.Println("Grasshopper is able to generate a Nuleculde Metadata snippet")
		fmt.Fprint(&buffer, generateNuleculeMetadata(NuleculeMetadata{labels["io.k8s.display-name"], versionString, labels["io.k8s.description"], ""}))
	}

	fmt.Fprintf(&buffer, `graph:
  - name: %s
`, generateNuleculeID(labels["io.k8s.display-name"]))

	// ok, lets see if we need to generate some requiremets for http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#storageRequirementsObject
	parameters := GetNuleculeParametersFromLabels(labels)
	if len(parameters) > 0 {
		jww.WARN.Println("Grasshopper will not generate a Nuleculde Parameters snippet with constraints!")
		jww.DEBUG.Println("Grasshopper is able to generate a Nuleculde Parameters snippet")

		fmt.Fprintln(&buffer, "    params:")

		for _, parameter := range parameters {
			fmt.Fprintf(&buffer, generateNuleculeParameter(parameter))
		}
	}

	volumes := GetNuleculeVolumesFromLabels(labels)

	// ja, looks like we need to...
	if len(volumes) > 0 {
		jww.DEBUG.Println("Grasshopper is able to generate a Nuleculde Requirements snippet")
		fmt.Fprint(&buffer, "requirements:\n")

		for _, volume := range volumes {
			fmt.Fprint(&buffer, generateNuleculePersistentVolume(volume))
		}
	}

	return buffer.String()
}

func generateNuleculeID(in string) string {
	return strings.Replace(in, " ", "_", -1)
}

func isRequired(maybe string) bool {
	return strings.Contains(maybe, "required")
}
