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
	"net/url"
	"strings"

	"github.com/goern/grasshopper/nulecule"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	jww "github.com/spf13/jwalterweatherman"
)

var dockerSchema string
var dockerHost string
var dockerPort int
var dockerPath string
var dockerTLSVerify bool
var dockerCertPath string
var dockerRegistry string

var fetchCmdV *cobra.Command

//FetchFunction is the function that downloads all Nulecule container images
func FetchFunction(cmd *cobra.Command, args []string) {
	InitializeConfig()

	if len(args) < 1 {
		cmd.Usage()
		jww.FATAL.Println("URL to be fetched is missing")
		return
	}

	url, err := url.Parse(args[0])
	if err != nil {
		jww.ERROR.Printf("%s is not a valid URL\n", args[0])
		return
	}

	// let's load the Nulecule of the application we want to deploy
	// this will recursively load all Nulecule files and merge them
	jww.INFO.Printf("fetching: %q", strings.Join(args, " "))
	_, err = nulecule.LoadNulecule(url) // TODO

	if err != nil {
		jww.ERROR.Printf("Can't load Nulecule from %s\n", args[0])
		return
	}
}

//FetchCmd returns an initialized CLI fetch command
var FetchCmd = &cobra.Command{
	Use:   "fetch URL",
	Short: "Download application from URL",
	Long:  `Will download an application from a URL and combine artifacts from the target application and any dependent applications.`,
	Run:   FetchFunction,
}

func init() {
	GrasshopperCmd.AddCommand(FetchCmd)

	FetchCmd.PersistentFlags().StringVar(&dockerHost, "docker-host", "localhost", "This is the host running the docker endpoint")
	FetchCmd.PersistentFlags().BoolVar(&dockerTLSVerify, "docker-tls-verify", false, "perform TLS certificate verification on connect")
	FetchCmd.PersistentFlags().StringVar(&dockerCertPath, "docker-cert-path", "/etc/docker/certs", "X.509 certificate path to be used during TLS certificate verification")
	FetchCmd.PersistentFlags().StringVar(&dockerRegistry, "docker-registrz", "registry.docker.com", "the registry to fetch Nulecules from")

	// default settings
	viper.SetDefault("dockerHost", "localhost")
	viper.SetDefault("dockerPath", "/var/run/docker.sock")
	viper.SetDefault("dockerSchema", "unix")
	viper.SetDefault("dockerCertPath", "/etc/docker/certs")
	viper.SetDefault("dockerTLSVerify", false)
	viper.SetDefault("dockerRegistry", "registry.docker.com")

	// bind config to command flags
	if FetchCmd.PersistentFlags().Lookup("docker-host").Changed {
		viper.Set("dockerHost", Verbose)
	}

}
