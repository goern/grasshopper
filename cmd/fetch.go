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
var dockerSocketPath string
var dockerTLSVerify bool
var dockerCertPath string
var dockerRegistry string

//FetchCmd returns an initialized CLI fetch command
var FetchCmd = &cobra.Command{
	Use:   "fetch URL",
	Short: "Download application from URL",
	Long:  `fetch will download an application from a URL and combine all components from the that application and any dependent applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
			jww.FATAL.Println("URL to be fetched is missing")
			return
		}

		if err := InitializeConfig(); err != nil {
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

		// TODO we need to check the command line!
		var dockerEndpoint = nulecule.DockerEndpoint{"unix", "", 0, "/var/run/docker.sock"}
		var loaderConfig = nulecule.LoaderOptions{viper.GetString("docker-registry"), &dockerEndpoint}

		jww.DEBUG.Printf("Nulecule Loader Options: %#v", loaderConfig)
		jww.DEBUG.Printf("Nulecule Loader Options: Endpoint: %#v", loaderConfig.Endpoint)

		_, err = nulecule.LoadNulecule(&loaderConfig, url)

		if err != nil {
			jww.ERROR.Printf("Can't load Nulecule from %s\n", args[0])
			return
		}
	},
}

func init() {
	FetchCmd.Flags().StringVar(&dockerHost, "docker-host", "localhost", "This is the host running the docker endpoint")
	FetchCmd.Flags().BoolVar(&dockerTLSVerify, "docker-tls-verify", false, "perform TLS certificate verification on connect")
	FetchCmd.Flags().StringVar(&dockerCertPath, "docker-cert-path", "/etc/docker/certs", "X.509 certificate path to be used during TLS certificate verification")
	FetchCmd.Flags().StringVar(&dockerRegistry, "docker-registry", "registry.docker.com", "the registry to fetch Nulecules from")

	viper.BindPFlag("docker-host", FetchCmd.Flags().Lookup("docker-host"))
	viper.BindPFlag("docker-tls-verify", FetchCmd.Flags().Lookup("docker-tls-verify"))
	viper.BindPFlag("docker-cert-path", FetchCmd.Flags().Lookup("docker-cert-path"))
	viper.BindPFlag("docker-registry", FetchCmd.Flags().Lookup("docker-registry"))

	// default settings
	viper.SetDefault("dockerHost", "localhost")
	viper.SetDefault("dockerSocketPath", "/var/run/docker.sock")
	viper.SetDefault("dockerSchema", "unix")
	viper.SetDefault("dockerCertPath", "/etc/docker/certs")
	viper.SetDefault("dockerTLSVerify", false)
	viper.SetDefault("dockerRegistry", "registry.docker.com")
}
