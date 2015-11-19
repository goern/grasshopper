// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package nulecule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	assert := assert.New(t)

	containerApplication, parseError := ParseFile("../test-fixtures/Nulecule")

	if parseError != nil {
		t.Log(parseError)
	}

	assert.NotNil(containerApplication)

	containerApplication, parseError = ParseFile("../test-fixtures/with-inherits")

	if parseError != nil {
		t.Log(parseError)
	}

	assert.NotNil(containerApplication)

	containerApplication, parseError = ParseFile("../test-fixtures/with-constraints")

	if parseError != nil {
		t.Log(parseError)
	}

	assert.NotNil(containerApplication)

	//	t.Log(string(containerApplication))
}
