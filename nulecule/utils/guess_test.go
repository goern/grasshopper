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
	const expectedPersistenVolume1GiB = `---
- persistentVolume:
    name: "var-lib-psql-data"
    accessMode: "ReadWrite"
    size: 1
`

	// This is a default config, using 4GiB
	const expectedPersistenVolume4GiBDefault = `---
- persistentVolume:
    name: "var-lib-psql-data"
    accessMode: "ReadWrite"
    size: 4 # GB by default
`

	// This is a claim for 4GiB
	const expectedPersistenVolume4GiB = `---
- persistentVolume:
    name: "var-lib-psql-data"
    accessMode: "ReadWrite"
    size: 4
`

	assert.Equal(expectedPersistenVolume1GiB, generateNuleculePersistentVolume(NuleculePersistentVolume{"/var/lib/psql/data", 1}))

	assert.Equal(expectedPersistenVolume4GiBDefault, generateNuleculePersistentVolume(NuleculePersistentVolume{"/var/lib/psql/data", -1}))

	assert.Equal(expectedPersistenVolume4GiBDefault, generateNuleculePersistentVolume(NuleculePersistentVolume{"/var/lib/psql/data", 0}))

	assert.Equal(expectedPersistenVolume4GiB, generateNuleculePersistentVolume(NuleculePersistentVolume{"/var/lib/psql/data", 4}))

}
