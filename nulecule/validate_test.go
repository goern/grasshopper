// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package nulecule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	assert := assert.New(t)

	containerApplication, parseError := ParseFile("../test-fixtures/Nulecule")

	if parseError != nil {
		t.Log(parseError)
	}

	if assert.NotNil(containerApplication) {
		assert.Equal(NuleculeVersion, containerApplication.Specversion, "Nulecule Spec Version should be 0.0.2")

		valErr := containerApplication.Validate()

		if valErr != nil {
			t.Log(valErr)
		}
	}

	containerApplicationBroken, parseError := ParseFile("../test-fixtures/Nulecule")

	if parseError != nil {
		t.Log(parseError)
	}

	if assert.NotNil(containerApplicationBroken) {

	}
}
