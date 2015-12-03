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
	"strings"

	"github.com/goern/grasshopper/nulecule"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

func doFetchFromURL(URL string) {
	// TODO get a from from a URL via "github.com/hashicorp/go-getter"
}

//FetchFunction is the function that downloads all Nulecule container images
func FetchFunction(cmd *cobra.Command, args []string) {
	if Verbose { // FIXME do we need this?
		jww.SetLogThreshold(jww.LevelTrace)
		jww.SetStdoutThreshold(jww.LevelInfo)
	}

	if len(args) < 1 {
		cmd.Usage()
		jww.FATAL.Println("URL to be fetched is missing")
	}

	jww.INFO.Printf("fetching: %q", strings.Join(args, " "))
	app, err := nulecule.LoadNulecule(args[0])

	if err != nil {
		jww.ERROR.Printf("Can't load Nulecule from %s\n", args[0])
		return
	}

	fmt.Println(app)
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

	FetchCmd.PersistentFlags().String("docker-host", "localhost", "This is the host running the docker endpoint")
	FetchCmd.PersistentFlags().Bool("docker-tls-verify", true, "perform TLS certificate verification on connect")
	FetchCmd.PersistentFlags().String("docker-cert-path", "/etc/docker/certs", "X.509 certificate path to be used during TLS certificate verification")

}
