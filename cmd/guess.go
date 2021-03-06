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

	"github.com/goern/grasshopper/nulecule/utils"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

//GuessCmd returns an initialized guess command
var GuessCmd = &cobra.Command{
	Use:   "guess DOCKERFILE",
	Short: "guess something from a DOCKERFILE",
	Long:  "Guess some information from a Dockerfile that can be handy in a Nulecule.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := InitializeConfig(); err != nil {
			return
		}

		// figure out if a Dockerfile was provided
		if len(args) < 1 {
			cmd.Usage()
			jww.FATAL.Println("path to a Dockerfile is required but not supplied")
			return
		}

		// if we got a Dockerfile
		jww.INFO.Println("guessing from " + string(args[0]))

		// lets parse it and guess some Nulecule file from it
		_, nuleculeFile, err := utils.GuessFromDockerfile(args[0])
		if err != nil {
			jww.FATAL.Println("failed to read the Dockerfile, nothing guessed")
			return
		}

		// it was good guessing that Nulecule? Print it!
		if nuleculeFile != "" {
			fmt.Println(nuleculeFile)
		}

	},
}
