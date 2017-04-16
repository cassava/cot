// Copyright (c) 2017, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

import (
	"errors"
	"os"

	"github.com/cassava/cot"
	"github.com/spf13/cobra"
)

func init() {
	MainCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [directory]",
	Short: "initialize a new cot repository",
	Long: `Initialize a new cot repository.

  If no directory is specified on the command line, then the current directory
  is used. Multiple directories may not be specified.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		n := len(args)
		var path string
		if n == 0 {
			d, err := os.Getwd()
			if err != nil {
				return err
			}
			path = d
		} else if n == 1 {
			path = args[0]
		} else {
			return errors.New("incorrect number of arguments")
		}
		return cot.Init(path)
	},
}
