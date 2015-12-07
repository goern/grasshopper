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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	jww "github.com/spf13/jwalterweatherman"
)

//GrasshopperCmd is Grasshopper's root command.
var GrasshopperCmd = &cobra.Command{
	Use:   "grasshopper",
	Short: "make a Nulecule GO!",
	Long:  `Grasshopper is a GOlang implementation of the Nulecule Specification.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(`Grasshopper  Copyright (C) 2015  Christoph Görn
This program comes with ABSOLUTELY NO WARRANTY; for details use 'grasshopper --help'.
This is free software, and you are welcome to redistribute it
under certain conditions; use 'grasshopper show license' for details.`)

		if err := InitializeConfig(); err != nil {
			return err
		}

		return nil
	},
}

var grasshopperCmdV *cobra.Command

//NuleculeCmd is Grasshopper's Nulecule sub-command.
var NuleculeCmd = &cobra.Command{
	Use:   "nulecule",
	Short: "do the nulecule",
	Long:  `Work with a Nulecule.`,
}

//versionCmd will simply print the version string of Grasshopper
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Grasshopper",
	Long:  "All software has versions. This is the Grasshopper's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Grasshopper %s (%s)\n", version, minversion)
	},
}

//bashAutogenerateCmd will write out some bach autocompletion shell script
// its unsupported by now, and not exposed to the command line binary
var bashAutogenerateCmd = &cobra.Command{
	Use:   "generate-bash-autocompletion",
	Short: "Generate an autocompletion bash shell script",
	Long:  "This will generate an autocompletion bash shell script in the current directory.",
	Run: func(cmd *cobra.Command, args []string) {
		GrasshopperCmd.GenBashCompletionFile("grasshopper_autocompletion.sh")
	},
}

//Verbose will enable more verbose logging
var Verbose bool

//Quiet is the opposite of Verbose
var Quiet bool

//DryRun will pretend to do something, but really really doesnt do anything
var DryRun bool

//Log will write to a logfile
var Log bool

//Experimental will enable experimental output
var Experimental bool

var version string    // set by -X via Makefile
var minversion string // set by -X via Makefile

//Execute adds all child commands to the root command GrasshopperCmd and sets flags appropriately.
func Execute() {
	GrasshopperCmd.SuggestionsMinimumDistance = 1

	//add child commands to the root command.
	GrasshopperCmd.AddCommand(versionCmd)

	// GrasshopperCmd.AddCommand(bashAutogenerateCmd)

	// add nulecule and it's sub-commands
	NuleculeCmd.AddCommand(IndexCmd)
	NuleculeCmd.AddCommand(GuessCmd)
	GrasshopperCmd.AddCommand(NuleculeCmd)

	GrasshopperCmd.AddCommand(FetchCmd)
	// GrasshopperCmd.AddCommand(InstallCmd)
	// GrasshopperCmd.AddCommand(RunCmd)
	// GrasshopperCmd.AddCommand(StopCmd)
	// GrasshopperCmd.AddCommand(UninstallCmd)
	// GrasshopperCmd.AddCommand(CleanCmd)

	if err := GrasshopperCmd.Execute(); err != nil {
		// the err is already logged by Cobra
		os.Exit(-1)
	}

}

func init() {
	GrasshopperCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	GrasshopperCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "quiet output")
	GrasshopperCmd.PersistentFlags().BoolVarP(&Log, "log", "l", true, "write logging output to file")
	GrasshopperCmd.PersistentFlags().BoolVarP(&Experimental, "experimental", "x", true, "write experimental output to stdout")

	grasshopperCmdV = GrasshopperCmd

	viper.BindPFlag("verbose", GrasshopperCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("quiet", GrasshopperCmd.PersistentFlags().Lookup("quiet"))
	viper.BindPFlag("log", GrasshopperCmd.PersistentFlags().Lookup("log"))
	viper.BindPFlag("experimental", GrasshopperCmd.PersistentFlags().Lookup("experimental"))

	if Log {
		jww.SetLogFile("grasshopper.log")
	}

	if Quiet {
		jww.SetStdoutThreshold(jww.LevelWarn)
	}

	if Verbose {
		jww.SetLogThreshold(jww.LevelTrace)
		jww.SetStdoutThreshold(jww.LevelTrace)
	}

}

// InitializeConfig reads in config file and ENV variables if set.
func InitializeConfig(subCmdVs ...*cobra.Command) error {
	viper.SetConfigType("json")
	viper.SetConfigName("grasshopper") // name of config file (without extension)
	//	viper.AddConfigPath("/etc/grasshopper.d/")  // path to look for the config file
	//	viper.AddConfigPath("$HOME/.grasshopper.d") // call multiple times to add many search paths
	viper.AddConfigPath(".") // optionally look for config in the working directory

	// read config from storage
	err := viper.ReadInConfig()
	if err != nil {
		jww.WARN.Printf("Unable to read Config file. %#v I will fall back to my defaults...", err)
		err = nil // we just skip this error
	}

	// set some sane defaults
	viper.SetDefault("Verbose", false)
	viper.SetDefault("Quiet", false)
	viper.SetDefault("Log", true)
	viper.SetDefault("Experimental", true)

	if grasshopperCmdV.PersistentFlags().Lookup("verbose").Changed {
		viper.Set("Verbose", Verbose)
	}

	if viper.GetBool("verbose") {
		jww.SetStdoutThreshold(jww.LevelTrace)
		jww.SetLogThreshold(jww.LevelTrace)
	}

	return err
}
