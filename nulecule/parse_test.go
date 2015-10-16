// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package nulecule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	assert := assert.New(t)

	containerApplication, parseError := ParseFile("../fixtures/Nulecule")

	if parseError != nil {
		t.Log(parseError)
	}

	if assert.NotNil(containerApplication) {
		assert.Equal(NuleculeVersion, containerApplication.Specversion, "Nulecule Spec Version should be 0.0.2")
	}

	//	t.Log(string(containerApplication))
}
