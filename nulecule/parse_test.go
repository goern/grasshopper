// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package nulecule

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("../fixtures/Nulecule")

	if err != nil {
		panic(err)
	}

	containerApplication, parseError := Parse(f)

	if parseError != nil {
		panic(parseError)
	}

	f.Close()

	t.Log(containerApplication)
}
