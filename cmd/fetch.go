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

package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

func doFetchFromURL(URL string) {
	// TODO get a from from a URL via "github.com/hashicorp/go-getter"
}

//FetchFunction is the function that downloads all Nulecule container images
func FetchFunction(cmd *cobra.Command, args []string) {
	if Verbose {
		jww.SetLogThreshold(jww.LevelTrace)
		jww.SetStdoutThreshold(jww.LevelInfo)
	}

	if len(args) < 1 {
		cmd.Usage()
		jww.FATAL.Println("URL to be fetched is missing")
	}

	jww.INFO.Printf("fetching: %q", strings.Join(args, " "))
}

//FetchCmd returns an initialized CLI fetch command
var FetchCmd = &cobra.Command{
	Use:   "fetch URL",
	Short: "Download application from URL",
	Long:  `Will download an application from a URL and combine artifacts from the target application and any dependent applications.`,
	Run:   FetchFunction,
}
