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

	"github.com/goern/grasshopper/nulecule"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xeipuuv/gojsonschema"

	jww "github.com/spf13/jwalterweatherman"
)

// SpecVersion specifies the Nulecule Specification's version to validate with
var SpecVersion string

// validateCmd respresents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate if a Nulecule file complies to a Specification",
	Long: `validate will validate that a Nulecule file complies to a specified version
of the Nulecule Specification. Currently version 0.0.2 is supported.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("validate called")

		if err := InitializeConfig(); err != nil {
			return
		}

		jww.DEBUG.Printf("Nulecule Validation with Specification Version: %s", SpecVersion)

		// check if SpecVersion equals nulecule.NuleculeReleasedVersions
		if SpecVersion != nulecule.NuleculeReleasedVersions {
			jww.ERROR.Printf("The specified version (%s) of the Nulecule Specification is invalid\n", SpecVersion)
		}

		schemaLoader := gojsonschema.NewReferenceLoader("https://raw.githubusercontent.com/projectatomic/nulecule/master/spec/schema.json")
		documentLoader := gojsonschema.NewReferenceLoader("file:///home/me/document.json")

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			panic(err.Error())
		}

		if result.Valid() {
			fmt.Printf("The document is valid\n")
		} else {
			fmt.Printf("The document is not valid. see errors :\n")
			for _, desc := range result.Errors() {
				fmt.Printf("- %s\n", desc)
			}
		}
	},
}

func init() {
	NuleculeCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringVarP(&SpecVersion, "spec-version", "s", "0.0.2", "Nulecule Specification version to use")
	viper.BindPFlag("spec-version", FetchCmd.Flags().Lookup("spec-version"))

	viper.SetDefault("spec-version", "0.0.2")

}
