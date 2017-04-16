// Copyright (c) 2017, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/cassava/cot"
	"github.com/cassava/cot/internal"
	"github.com/goulash/color"
	"github.com/spf13/cobra"
)

// Set up logging.
func init() {
	cot.Logger().Formatter = internal.ConsoleFormatter{}
	cot.Logger().Level = logrus.InfoLevel
}

// col lets us print in colors
var col = color.New()

var (
	quiet      bool
	verbose    bool
	colorState string
)

func init() {
	MainCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	MainCmd.PersistentFlags().BoolVarP(&verbose, "quiet", "q", false, "quiet output, supersedes verbose flag")
	MainCmd.PersistentFlags().StringVar(&colorState, "color", "auto", "whether to use color (always|auto|never)")
}

var MainCmd = &cobra.Command{
	Use:   "cot [options] <command>",
	Short: "Cot is a tool to administer your dotfiles.",
	Long: `Cot is a tool to administer your dotfiles.
Built with much love by cassava in Go.

Complete documentation is available at http://github.com/cassava/cot
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// This function can be overriden if it's not necessary for a command.
		cmd.SilenceErrors = true
		cmd.SilenceUsage = true

		if quiet {
			cot.Logger().Level = logrus.WarnLevel
		} else if verbose {
			cot.Logger().Level = logrus.DebugLevel
		}

		if colorState == "auto" {
			col.SetFile(os.Stdout)
		} else if colorState == "always" {
			col.SetEnabled(true)
		} else if colorState == "never" {
			col.SetEnabled(false)
		}

		return nil
	},
}

func main() {
	if err := MainCmd.Execute(); err != nil {
		col.Fprintf(os.Stderr, "@rError: %s\n", err)
	}
}
