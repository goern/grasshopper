// Copyright ©2015 NAME HERE <EMAIL ADDRESS>
//
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package cmd

import "github.com/spf13/cobra"

// showCmd respresents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show a few things",
	Long:  `This will show you a few things that Grasshopper can do for you.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
			return
		}
	},
}

func init() {
	GrasshopperCmd.AddCommand(showCmd)
}
