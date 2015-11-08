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
	"github.com/op/go-logging"
	"github.com/spf13/cobra"
	//  "github.com/hashicorp/go-getter"
)

var log = logging.MustGetLogger("grasshopper")

//dryrun will force the fetchCmd to not do anything at all
var dryrun bool

//FetchFunction is the function that downloads all Nulecule container images
func FetchFunction(cmd *cobra.Command, args []string) {
	log.Info("fetching: ", args[0])
}

//FetchCmd returns an initialized CLI fetch command
var FetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download application from URL",
	Long:  `Will download from a URL and combine artifacts from the target application and any dependent applications.`,
	Run:   FetchFunction,
}

// FetchCmd.Flags().StringVarP(&dryrun, "dry-run", "d", "", "do not really do anything")
