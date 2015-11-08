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

// Package main is the main command line tool for Grasshopper.
package main

import (
	"fmt"

	"github.com/op/go-logging"
	"github.com/spf13/cobra"

	"github.com/goern/grasshopper/cmd"
)

var version string
var minversion string
var log = logging.MustGetLogger("grasshopper")

//Verbose is a global --verbose command line thingy
var Verbose bool

func main() {
	var GrasshopperCmd = &cobra.Command{
		Use:   "grasshopper",
		Short: "make a Nulecule GO!",
		Long:  `Grasshopper is a GOlang implementation of the Nulecule Specification.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Grasshopper",
		Long:  `All software has versions. This is the Grasshopper's`,
		Run: func(cmd *cobra.Command, args []string) {
			versionString := ""

			if Verbose {
				versionString = fmt.Sprintf("Grasshopper v%s (%s)", version, minversion)
			} else {
				versionString = fmt.Sprintf("Grasshopper v%s", version)
			}

			fmt.Println(versionString)
		},
	}

	var bashAutogenerateCmd = &cobra.Command{
		Use:   "generate-bash-autocompletion",
		Short: "Generate an autocompletion bash shell script",
		Long:  "This will generate an autocompletion bash shell script in the current directory.",
		Run: func(cmd *cobra.Command, args []string) {
			GrasshopperCmd.GenBashCompletionFile("grasshopper_autocompletion.sh")
		},
	}

	GrasshopperCmd.SuggestionsMinimumDistance = 1

	GrasshopperCmd.AddCommand(versionCmd)
	GrasshopperCmd.AddCommand(bashAutogenerateCmd)
	GrasshopperCmd.AddCommand(cmd.FetchCmd)
	GrasshopperCmd.AddCommand(cmd.InstallCmd)
	GrasshopperCmd.AddCommand(cmd.RunCmd)
	GrasshopperCmd.AddCommand(cmd.StopCmd)
	GrasshopperCmd.AddCommand(cmd.UninstallCmd)
	GrasshopperCmd.AddCommand(cmd.CleanCmd)

	GrasshopperCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	GrasshopperCmd.Execute()

	/*

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
	*/
}
