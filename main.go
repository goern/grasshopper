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

	"github.com/goern/grasshopper/cmd"
	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var version string
var minversion string
var log = logging.MustGetLogger("grasshopper")

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

			versionString = fmt.Sprintf("Grasshopper v%s (%s)", version, minversion)

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
	//	cmd.FetchCmd.Flags().BoolVarP(&grasshopper.DryRun, "dry-run", "d", false, "dry run, just pretend to do something")

	GrasshopperCmd.AddCommand(cmd.InstallCmd)
	//	cmd.InstallCmd.Flags().StringVarP(&grasshopper.Provider, "provider", "p", "kubernetes", "Provider to be used, it may be 'kubernetes', 'openshift' or 'docker'")

	GrasshopperCmd.AddCommand(cmd.RunCmd)
	//	cmd.RunCmd.Flags().StringVarP(&grasshopper.Provider, "provider", "p", "kubernetes", "Provider to be used, it may be 'kubernetes', 'openshift' or 'docker'")

	GrasshopperCmd.AddCommand(cmd.StopCmd)
	//	cmd.StopCmd.Flags().StringVarP(&grasshopper.Provider, "provider", "p", "kubernetes", "Provider to be used, it may be 'kubernetes', 'openshift' or 'docker'")

	GrasshopperCmd.AddCommand(cmd.UninstallCmd)
	//	cmd.UninstallCmd.Flags().StringVarP(&grasshopper.Provider, "provider", "p", "kubernetes", "Provider to be used, it may be 'kubernetes', 'openshift' or 'docker'")

	GrasshopperCmd.AddCommand(cmd.CleanCmd)

	//	GrasshopperCmd.PersistentFlags().BoolVarP(&grasshopper.Verbose, "verbose", "v", false, "verbose output")

	GrasshopperCmd.Execute()

}
