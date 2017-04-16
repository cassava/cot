// Copyright (c) 2017, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

func init() {
	MainCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:               "version",
	Short:             "show version and date information",
	Long:              "Show the official version number of cot, as well as the release date.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		var progInfo = struct {
			Name      string
			Author    string
			Email     string
			Version   string
			Date      string
			Homepage  string
			Copyright string
			License   string
		}{
			Name:      "cot",
			Author:    "Ben Morgan",
			Email:     "neembi@gmail.com",
			Version:   "0.1",
			Date:      "14 April, 2017",
			Copyright: "2017",
			Homepage:  "https://github.com/cassava/cot",
			License:   "MIT",
		}
		versionTmpl.Execute(os.Stdout, progInfo)
	},
}

var versionTmpl = template.Must(template.New("version").Parse(`{{.Name}} version {{.Version}} ({{.Date}})
Copyright {{.Copyright}}, {{.Author}} <{{.Email}}>

You may find {{.Name}} on the Internet at
    {{.Homepage}}
Please report any bugs you may encounter.

The source code of {{.Name}} is licensed under the {{.License}} license.
`))
