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
	"archive/zip"
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"

	"github.com/goern/grasshopper/nulecule"
)

func init() {
	IndexCmd.AddCommand(indexListCmd)
	IndexCmd.AddCommand(indexInfoCmd)

}

//IndexCommand will interact with the Nulecule Library on github
var IndexCmd = &cobra.Command{
	Use:   "index",
	Short: "list index of applications or get info on one application",
	Long: `list all Nulecules or the details of one Nulecule on the Nulecule Library

index requires a subcommand, e.g. ` + "`grasshopper index list`.",
	Run: nil,
}

var indexListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all applications",
	Long:  `List all applications in the Nulecule Library Index.`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()

		if Verbose {
			jww.SetLogThreshold(jww.LevelTrace)
			jww.SetStdoutThreshold(jww.LevelInfo)
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		res, err := client.Get("https://github.com/projectatomic/nulecule-library/archive/master.zip")
		if err != nil {
			jww.FATAL.Printf("Fetching Nulecule Library zip URL failed with error %q\n", err)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			jww.FATAL.Printf("Failed to read response body with error %q\n", err)
		}

		r := bytes.NewReader(b)
		nuleculeLibraryIndexZip, err := zip.NewReader(r, int64(r.Len()))

		// Iterate through the files in the archive
		for _, item := range nuleculeLibraryIndexZip.File {
			if item.FileInfo().IsDir() {
				continue
			}

			// if the file is a Nulecule, getit!
			if item.FileInfo().Name() == "Nulecule" {
				jww.DEBUG.Printf("Found a Nulecule, size of it's description is %d\n", item.FileInfo().Size())

				rc, err := item.Open()
				if err != nil {
					jww.FATAL.Println(err)
					return
				}
				defer rc.Close()

				// get the Nulecules content
				nuci, parseError := nulecule.Parse(rc)

				if parseError != nil {
					jww.INFO.Println(parseError, " This may be due to unsupported (by Grasshopper) artifact inheritance.")
					continue
				}

				jww.INFO.Printf("%v\n", nuci)
			}

		}

		defer res.Body.Close()
	},
}

var indexInfoCmd = &cobra.Command{
	Use:   "info APPNAME",
	Short: "show info of APPNAME",
	Long:  `Show detailed info of APPNAME in the Nulecule Library Index.`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()

		// TODO require one argument
	},
}
