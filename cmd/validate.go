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
	"strings"

	"github.com/goern/grasshopper/nulecule"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	jww "github.com/spf13/jwalterweatherman"
)

// SpecVersion specifies the Nulecule Specification's version to validate with
var SpecVersion string

// validateCmd respresents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate URL",
	Short: "Validate if a Nulecule is compliant with the Specification",
	Long: `validate will validate that a Nulecule file at the supplied URL complies
to a specified version of the Nulecule Specification. The Nulecule file must be
in JSON Format.

If a docker:// URL to a Nulecule is provided, validate will also verify that
all artifacts are present within the Nulecule and that all artifacts are
parsable (JSON or YAML).

Currently version 0.0.2 of the Nulecule Specification is supported.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := InitializeConfig(); err != nil {
			return
		}

		if len(args) < 1 {
			cmd.Usage()
			jww.FATAL.Println("URL of Nulecule to validate is missing")
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

		// TODO if location started with docker:// pull the Nulecule,
		// extract the Nulecule file and validate it, do this recursively
		if strings.HasPrefix(location.String(), "docker://") {
			/*
				// check if the Nulecule's content itself is valid
				// TODO we need to check the command line!
				var dockerEndpoint = nulecule.DockerEndpoint{"unix", "", 0, "/var/run/docker.sock"}
				var loaderConfig = nulecule.LoaderOptions{viper.GetString("docker-registry"), &dockerEndpoint}
				app, err := nulecule.LoadNulecule(&loaderConfig, location)

				if err != nil {
					jww.ERROR.Printf("Can't load Nulecule from %s\n", location)
					return
				}

				err = app.Validate()
				if err != nil {
					jww.ERROR.Printf("Can't validate Nulecule's content: %#v\n", err)
					return
				}

				fmt.Printf("Nulecule's content is a valid Nulecule v%s\n", location.String(), SpecVersion)
			*/
		} else if strings.HasPrefix(location.String(), "file://") {

			// check if the Nulecule is valid JSON
			result, err := nulecule.ValidateSchema(SpecVersion, location)
			if err != nil {
				jww.ERROR.Printf("Can't validate Nulecule file: %#v\n", err)
				return
			}

			if result {
				fmt.Printf("Nulecule file (%s) is compliant with Nulecule Specification v%s\n", location.String(), SpecVersion)
				return
			}

		}

		fmt.Printf("The Nulecule file (%s) is not valid.\n", args[0])
	},
}

func init() {
	NuleculeCmd.AddCommand(validateCmd)

	/*
		validateCmd.Flags().StringVar(&dockerHost, "docker-host", "localhost", "This is the host running the docker endpoint")
		validateCmd.Flags().BoolVar(&dockerTLSVerify, "docker-tls-verify", false, "perform TLS certificate verification on connect")
		validateCmd.Flags().StringVar(&dockerCertPath, "docker-cert-path", "/etc/docker/certs", "X.509 certificate path to be used during TLS certificate verification")
		validateCmd.Flags().StringVar(&dockerRegistry, "docker-registry", "registry.docker.com", "the registry to fetch Nulecules from")

		viper.BindPFlag("docker-host", validateCmd.Flags().Lookup("docker-host"))
		viper.BindPFlag("docker-tls-verify", validateCmd.Flags().Lookup("docker-tls-verify"))
		viper.BindPFlag("docker-cert-path", validateCmd.Flags().Lookup("docker-cert-path"))
		viper.BindPFlag("docker-registry", validateCmd.Flags().Lookup("docker-registry"))
	*/

	validateCmd.Flags().StringVarP(&SpecVersion, "spec-version", "s", "0.0.2", "Nulecule Specification version to use")
	viper.BindPFlag("spec-version", FetchCmd.Flags().Lookup("spec-version"))

	viper.SetDefault("spec-version", "0.0.2")

}
