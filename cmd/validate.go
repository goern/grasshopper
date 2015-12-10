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

//Package cmd includes all the commands used for the grasshopper command. ;)
package cmd

import (
	"fmt"
	"net/url"

	"github.com/goern/grasshopper/nulecule"
	"github.com/goern/grasshopper/nulecule/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	jww "github.com/spf13/jwalterweatherman"
)

// SpecVersion specifies the Nulecule Specification's version to validate with
var SpecVersion string

// validateCmd respresents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate URL",
	Short: "Validate if a Nulecule file complies to a Specification",
	Long: `validate will validate that a Nulecule file at the supplied URL complies
to a specified version of the Nulecule Specification. Currently version 0.0.2
is supported.

The Nulecule file must be in JSON Format.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := InitializeConfig(); err != nil {
			return
		}

		if len(args) < 1 {
			cmd.Usage()
			jww.FATAL.Println("URL of Nulecule file to validate is missing")
			return
		}

		location, err := url.Parse(args[0])
		if err != nil {
			jww.ERROR.Printf("%s\n", err)
			return
		}

		jww.DEBUG.Printf("Nulecule Validation with Specification Version: %s", SpecVersion)

		// check if SpecVersion equals nulecule.NuleculeReleasedVersions
		if SpecVersion != nulecule.NuleculeReleasedVersions {
			jww.ERROR.Printf("The specified version (%s) of the Nulecule Specification is not allowed, it must be '0.0.2'\n", SpecVersion)
			return
		}

		// check if the Nulecule is JSON

		result, err := utils.Validate(SpecVersion, location)
		if err != nil {
			jww.ERROR.Printf("Can't validate Nulecule file: %#v\n", err)
			return
		}

		if result {
			fmt.Printf("Nulecule file (%s) complies to Nulecule Specification v%s\n", location.String(), SpecVersion)
			return
		}

		fmt.Printf("The Nulecule file (%s) is not valid.\n", args[0])
	},
}

func init() {
	NuleculeCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringVarP(&SpecVersion, "spec-version", "s", "0.0.2", "Nulecule Specification version to use")
	viper.BindPFlag("spec-version", FetchCmd.Flags().Lookup("spec-version"))

	viper.SetDefault("spec-version", "0.0.2")

}
