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

import "github.com/codegangsta/cli"

//RunCommand returns an initialized CLI run command
func RunCommand() cli.Command {
	return cli.Command{
		Name:  "run",
		Usage: "Run an application.",
		Action: func(c *cli.Context) {
			println("running: ", c.Args().First())
		},
		Flags: LifecycleFlagSet(),
	}
}
