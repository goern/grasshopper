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
	"github.com/codegangsta/cli"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("grasshopper")

func fetchFlagSet() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "dry-run",
			Usage: "This will cause fetch to not do anything, nor fetch something.",
		},
	}
}

//fetchFunction is the function that downloads all Nulecule container images
func fetchFunction(c *cli.Context) {
	if len(c.Args()) < 1 {
		cli.ShowCommandHelp(c, "fetch")
		log.Critical("Please provide a repository to install from.")
	} else {
		log.Info("fetching: ", c.Args().First())
	}
}

//FetchCommand returns an initialized CLI fetch command
func FetchCommand() cli.Command {
	return cli.Command{
		Name:   "fetch",
		Usage:  "Will download and combine artifacts from the target application and any dependent applications.",
		Action: fetchFunction,
		Flags:  fetchFlagSet(),
	}
}
