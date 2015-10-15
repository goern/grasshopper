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

//LifecycleFlagSet is applied to all Commands that operate on a provider and controll the lifecycle of an application within a provider's context
func LifecycleFlagSet() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "provider",
			Usage:  "The provider to use. Overrides provider value in answerfile.",
			EnvVar: "GRASSHOPPER_PROVIDER",
		},
	}
}

//InstallCommand returns an initialized CLI install command
func InstallCommand() cli.Command {
	return cli.Command{
		Name:  "install",
		Usage: "Peform actions to prepare application to run.",
		Action: func(c *cli.Context) {
			println("installing: ", c.Args().First())
		},
		Flags: LifecycleFlagSet(),
	}
}
