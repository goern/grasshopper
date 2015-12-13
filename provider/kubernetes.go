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

package provider

import (
	"fmt"
	"log"

	"gopkg.in/jmcvetta/napping.v3"
)

type ResponseUserAgent struct {
	Useragent string `json:"user-agent"`
}

func doGet() {
	e := struct {
		Message string
	}{}

	s := napping.Session{}
	res := ResponseUserAgent{}
	resp, err := s.Get("https://api.github.com/repos/goern/grasshopper", nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response Status:", resp.Status())

	if resp.Status() == 200 {
		fmt.Printf("Result: %s\n\n", resp.RawText())
		fmt.Println("res:", res.Useragent)
	} else {
		fmt.Println("Bad response status from server")
		fmt.Printf("\t Status:  %v\n", resp.Status())
		fmt.Printf("\t Message: %v\n", e.Message)
	}
}
