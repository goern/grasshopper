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

package main

import (
	"os"
	"time"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "grasshopper"
	app.Version = "0.0.2"
	app.Compiled = time.Now()
	app.Usage = "make a Nulecule go!"
	app.Authors = []cli.Author{
		{
			Name:  "Christoph Görn",
			Email: "goern@redhat.com",
		},
	}

	app.EnableBashCompletion = true

	// global level flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show more output",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "fetch",
			Usage: "Will download and combine artifacts from the target application and any dependent applications.",
			Action: func(c *cli.Context) {
				println("fetching: ", c.Args().First())
			},
		},
		{
			Name:  "install",
			Usage: "Peform actions to prepare application to run.",
			Action: func(c *cli.Context) {
				println("installing: ", c.Args().First())
			},
		},
		{
			Name:  "run",
			Usage: "Run an application.",
			Action: func(c *cli.Context) {
				println("running: ", c.Args().First())
			},
		},
		{
			Name:  "stop",
			Usage: "Stop an application.",
			Action: func(c *cli.Context) {
				println("stopping: ", c.Args().First())
			},
		},
		{
			Name:  "uninstall",
			Usage: "Remove deployment configuration from platform.",
			Action: func(c *cli.Context) {
				println("uninstalling: ", c.Args().First())
			},
		},
		{
			Name:  "clean",
			Usage: "Remove artifacts files from local system and clean up directory.",
			Action: func(c *cli.Context) {
				println("cleaning: ", c.Args().First())
			},
		},
	}

	app.Action = func(c *cli.Context) {
		println("GO! I say!")
	}

	app.Run(os.Args)
}
