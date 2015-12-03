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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Grasshopper  Copyright (C) 2015  Christoph Görn
This program comes with ABSOLUTELY NO WARRANTY; for details use 'grasshopper --help'.
This is free software, and you are welcome to redistribute it
under certain conditions; use 'grasshopper show license' for details.`)

		InitializeConfig()
	},
}

//NuleculeCmd is Grasshopper's Nulecule sub-command.
var NuleculeCmd = &cobra.Command{
	Use:   "nulecule",
	Short: "do the nulecule",
	Long:  `Work with a Nulecule.`,
}
var grasshopperCmdV *cobra.Command

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

//DoLog will write to a temporary logfile
var DoLog bool

//Experimental will enable experimental output
var Experimental bool

var version string    // set by -X via Makefile
var minversion string // set by -X via Makefile

//Execute adds all child commands to the root command GrasshopperCmd and sets flags appropriately.
func Execute() {
	jww.SetStdoutThreshold(jww.LevelInfo)

	GrasshopperCmd.SuggestionsMinimumDistance = 1

	//add child commands to the root command.
	GrasshopperCmd.AddCommand(versionCmd)

	GrasshopperCmd.AddCommand(bashAutogenerateCmd)

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

	/* FIXME this is nice, but we need a better one!
	manHeader := &cobra.GenManHeader{
		Title:   "grasshopper",
		Section: "1",
	}
	out := new(bytes.Buffer)
	GrasshopperCmd.GenMan(manHeader, out)
	fmt.Println(out.String())
	*/

	/* FIxME this is nice, but we need a better one!
	manHeader := &cobra.GenManHeader{
		Title:   "grasshopper",
		Section: "1",
	}
	out := new(bytes.Buffer)
	GrasshopperCmd.GenMan(manHeader, out)
	fmt.Println(out.String())
	*/

	if err := GrasshopperCmd.Execute(); err != nil {
		// the err is already logged by Cobra
		os.Exit(-1)
	}

}

// InitializeConfig initializes a config file with sensible default configuration flags.
func InitializeConfig() {
	if Quiet {
		jww.SetStdoutThreshold(jww.LevelWarn)
	}

	viper.SetConfigName("grasshopper")          // name of config file (without extension)
	viper.AddConfigPath("/etc/grasshopper.d/")  // path to look for the config file
	viper.AddConfigPath("$HOME/.grasshopper.d") // call multiple times to add many search paths
	viper.AddConfigPath(".")                    // optionally look for config in the working directory

	// read config from storage
	err := viper.ReadInConfig() // FIXME
	if err != nil {
		jww.WARN.Println("Unable to locate Config file. I will fall back to my defaults...")
	}

	// default settings
	viper.SetDefault("Verbose", false)
	viper.SetDefault("Quiet", false)
	viper.SetDefault("DryRun", false)
	viper.SetDefault("DoLog", true)
	viper.SetDefault("Experimental", true)

	// bind config to command flags
	if grasshopperCmdV.PersistentFlags().Lookup("verbose").Changed {
		viper.Set("Verbose", Verbose)
	}
	if grasshopperCmdV.PersistentFlags().Lookup("quiet").Changed {
		viper.Set("Quiet", Quiet)
	}
	if grasshopperCmdV.PersistentFlags().Lookup("log").Changed {
		viper.Set("DoLog", DoLog)
	}
	if grasshopperCmdV.PersistentFlags().Lookup("experimental").Changed {
		viper.Set("Experimental", Experimental)
	}

	if DoLog {
		jww.SetLogFile("grasshopper.log")
	}

	if Verbose {
		jww.SetLogThreshold(jww.LevelTrace)
		jww.SetStdoutThreshold(jww.LevelTrace)
	}
}

//Initializes flags
func init() {
	GrasshopperCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	GrasshopperCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "quiet output")
	GrasshopperCmd.PersistentFlags().BoolVarP(&DoLog, "log", "l", true, "write logging output to file")
	GrasshopperCmd.PersistentFlags().BoolVarP(&Experimental, "experimental", "x", true, "write experimental output to stdout")

	grasshopperCmdV = GrasshopperCmd

	FetchCmd.Flags().BoolVarP(&DryRun, "dry-run", "y", false, "dry run the fetch operation")
}
