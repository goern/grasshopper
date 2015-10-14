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
	"github.com/goern/grasshopper/cmd"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("grasshopper")

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
		cmd.FetchCommand(),
		cmd.InstallCommand(),
		cmd.RunCommand(),
		cmd.StopCommand(),
		cmd.UninstallCommand(),
		cmd.CleanCommand(),
	}

	app.Action = func(c *cli.Context) {
		println("GO! I say!")
	}

	app.Run(os.Args)
}
