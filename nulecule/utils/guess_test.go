// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateNuleculePersistentVolume(t *testing.T) {
	assert := assert.New(t)

	// This is for 1GiB
	const expectedPersistenVolume1GiB = `- persistentVolume:
    name: "var-lib-psql-data"
    accessMode: "ReadWrite"
    size: 1
`

	// This is a default config, using 4GiB
	const expectedPersistenVolume4GiBDefault = `- persistentVolume:
    name: "var-lib-psql-data"
    accessMode: "ReadWrite"
    size: 4 # GB by default
`

	// This is a claim for 4GiB
	const expectedPersistenVolume4GiB = `- persistentVolume:
    name: "var-lib-psql-data"
    accessMode: "ReadWrite"
    size: 4
`

	assert.Equal(expectedPersistenVolume1GiB, generateNuleculePersistentVolume(NuleculePersistentVolume{"/var/lib/psql/data", 1}))

	assert.Equal(expectedPersistenVolume4GiBDefault, generateNuleculePersistentVolume(NuleculePersistentVolume{"/var/lib/psql/data", -1}))

	assert.Equal(expectedPersistenVolume4GiBDefault, generateNuleculePersistentVolume(NuleculePersistentVolume{"/var/lib/psql/data", 0}))

	assert.Equal(expectedPersistenVolume4GiB, generateNuleculePersistentVolume(NuleculePersistentVolume{"/var/lib/psql/data", 4}))

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

	correctAnswer := map[string]string{
		"DATA": "want that one too",
		"LOGS": "want this it",
	}

	assert.Equal(correctAnswer, GetNuleculeVolumesFromLabels(labels))
}
