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
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

//UninstallCommand returns an initialized CLI install command
var UninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall application",
	Long:  "Remove deployment configuration from platform.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uninstalling: " + strings.Join(args, " "))
	},
}
