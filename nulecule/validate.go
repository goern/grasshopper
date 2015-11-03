// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package nulecule

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
)

// Validate validates the Nulecule file
func (nulecule *ContainerApplication) Validate() error {
	var result error

	// lets see if we are supposed to work on a 0.0.2 Nulecule
	if nulecule.Specversion != "0.0.2" {
		result = multierror.Append(result, fmt.Errorf(
			"'specversion' MUST be 0.0.2"))
	}

	return result
}
