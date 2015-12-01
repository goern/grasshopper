// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package utils

import (
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/goern/grasshopper/nulecule"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGenerateNuleculePersistentVolume(t *testing.T) {
	assert := assert.New(t)

	// This is for 1GiB
	const expectedPersistenVolume1GiB = `  - persistentVolume:
      name: "data"
      accessMode: "ReadWrite"
      size: "1Gi"
`

	// This is a default config, using 4GiB
	const expectedPersistenVolume4GiBDefault = `  - persistentVolume:
      name: "data"
      accessMode: "ReadWrite"
      size: "4Gi" # GB by default
`

	// This is a claim for 4GiB
	const expectedPersistenVolume4GiB = `  - persistentVolume:
      name: "data"
      accessMode: "ReadWrite"
      size: "4Gi"
`

	// This has a dash
	const expectedPersistenVolume2GiB = `  - persistentVolume:
      name: "data-2"
      accessMode: "ReadWrite"
      size: "2Gi"
`

	// This has two dashes
	const expectedPersistenVolume3GiB = `  - persistentVolume:
      name: "3-data-3"
      accessMode: "ReadWrite"
      size: "3Gi"
`

	// This has a slashes
	const expectedPersistenVolume5GiB = `  - persistentVolume:
      name: "name"
      accessMode: "ReadWrite"
      size: "5Gi"
`

	// This has two slashes
	const expectedPersistenVolume6GiB = `  - persistentVolume:
      name: "org_name"
      accessMode: "ReadWrite"
      size: "6Gi"
`

	assert.Equal(expectedPersistenVolume1GiB, generateNuleculePersistentVolume(NuleculePersistentVolume{"data", "1Gi"}))

	assert.Equal(expectedPersistenVolume4GiBDefault, generateNuleculePersistentVolume(NuleculePersistentVolume{"data", ""}))

	assert.Equal(expectedPersistenVolume4GiB, generateNuleculePersistentVolume(NuleculePersistentVolume{"data", "4Gi"}))

	assert.Equal(expectedPersistenVolume2GiB, generateNuleculePersistentVolume(NuleculePersistentVolume{"data-2", "2Gi"}))

	assert.Equal(expectedPersistenVolume3GiB, generateNuleculePersistentVolume(NuleculePersistentVolume{"3-data-3", "3Gi"}))

	assert.Equal(expectedPersistenVolume5GiB, generateNuleculePersistentVolume(NuleculePersistentVolume{"name", "5Gi"}))

	assert.Equal(expectedPersistenVolume6GiB, generateNuleculePersistentVolume(NuleculePersistentVolume{"org_name", "6Gi"}))
}

func TestInLabels(t *testing.T) {
	assert := assert.New(t)

	labels := map[string]string{
		"a":   "a value",
		"b":   "another value",
		"A.b": "a and b value",
	}

	assert.Equal(true, InLabels("a", labels))
	assert.Equal(true, InLabels("A", labels))
	assert.Equal(true, InLabels("a.b", labels))
	assert.Equal(false, InLabels("z", labels))
}

func TestGetNuleculeVolumesFromLabels(t *testing.T) {
	assert := assert.New(t)

	labels := map[string]string{
		"io.projectatomic.nulecule.volume.data":          "want that one too",
		"io.projectatomic.nulecule.volume.logs":          "want this it",
		"io.projectatomic.nulecule.volume":               "this is broken",
		"io.projectatomic.nulecule.environment.required": "to high level",
		"io.k8s.display-name":                            "dont want it",
	}

	correctAnswer := []NuleculePersistentVolume{
		NuleculePersistentVolume{Name: "data", Size: "4Gi"}, NuleculePersistentVolume{Name: "logs", Size: "4Gi"},
	}

	for _, answer := range correctAnswer {
		assert.Contains(GetNuleculeVolumesFromLabels(labels), answer)
	}

}

func TestGuessFromDockerfile(t *testing.T) {
	assert := assert.New(t)

	viper.Set("Experimental", true) // simulate --experimental command line flag

	_, nuleculeStruct, err := GuessFromDockerfile("../../test-fixtures/Dockerfile-postgresql.rhel7")
	// test if it did guess without any error
	assert.Nil(err)
	assert.NotNil(nuleculeStruct)

	// lets see if we got all the content we want
	assert.Contains(nuleculeStruct, "specversion: 0.0.2")
	assert.Contains(nuleculeStruct, "  appversion: \"9.4-1\"")
	assert.Contains(nuleculeStruct, "    name: \"data\"")
	assert.Contains(nuleculeStruct, "    size: \"1Gi\"")
	assert.Contains(nuleculeStruct, "    name: \"logs\"")
	assert.Contains(nuleculeStruct, "    size: \"512Mi\"")

	// and parse the yaml, just to be sure
	app := nulecule.ContainerApplication{}
	unmarschalError := yaml.Unmarshal([]byte(nuleculeStruct), &app)
	assert.Nil(unmarschalError)
}

func TestSnippetsFromLabelsMap(t *testing.T) {
	assert := assert.New(t)

	labels := map[string]string{
		"io.k8s.description":  "This is a map of labels used for testing",
		"io.k8s.display-name": "Test Case TestSnippetsFromLabelsMap",
		"Version":             "1.0",
	}

	result := snippetsFromLabelsMap(labels)
	assert.NotNil(result)

	assert.Contains(result, "1.0")
}

func TestGetNuleculeParametersFromLabels(t *testing.T) {
	assert := assert.New(t)

	labels := map[string]string{
		"io.projectatomic.nulecule.environment.required": "POSTGRESQL_USER, POSTGRESQL_PASSWORD, POSTGRESQL_DATABASE",
		"io.projectatomic.nulecule.environment.optional": "POSTGRESQL_ADMIN_PASSWORD",
	}

	correctAnswer := []nulecule.Param{
		nulecule.Param{"POSTGRESQL_USER", "", nulecule.Constraint{}, "", true},
		nulecule.Param{"POSTGRESQL_PASSWORD", "", nulecule.Constraint{}, "", true},
		nulecule.Param{"POSTGRESQL_DATABASE", "", nulecule.Constraint{}, "", true},
		nulecule.Param{"POSTGRESQL_ADMIN_PASSWORD", "", nulecule.Constraint{}, "", false},
	}

	for _, answer := range correctAnswer {
		assert.Contains(GetNuleculeParametersFromLabels(labels), answer)
	}
}
