// Copyright (c) 2017, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package cot

import (
	"errors"
	"os"
	"path/filepath"
	"text/template"

	"github.com/cassava/cot/pathutil"
	"github.com/goulash/osutil"
)

var (
	CotDir          = []string{".cot/"}
	CotDirConfig    = []string{"config", "config.toml"}
	CotDirExclude   = []string{"info/exclude"}
	CotDirDirectory = []string{"info/directory"}
	CotDirHooks     = []string{"hooks/"}

	CotConfig  = []string{".cotconfig", ".cotconfig.toml"}
	CotExclude = []string{".cotignore"}
)

var shortConfigTmpl = template.Must(template.New("cot-short-config").Parse(
	`# cot repository configuration
name = {{.Name}}
exclude = ['.*']
directory = ['.local', '.local/*', '.config'
`))

var configTmpl = template.Must(template.New("cot-config").Parse(
	`# cot repository configuration
name = {{.Name}}
`))

var excludeTmpl = template.Must(template.New("cot-exclude").Parse(
	`# cot ls --exclude-from=.cot/info/exclude
# Lines that start with '#' are comments.
# The following source files are ignored.
.*
`))

var directoryTmpl = template.Must(template.New("cot-directory").Parse(
	`# cot ls --directory-from=.cot/info/directory
# Lines that start with '#' are comments.
# The following _target directories_ are be created, not linked.
.local
.local/*
.config
`))

var hooksReadMeTmpl = template.Must(template.New("cot-hooks-readme").Parse(
	`Cot hooks README

You can write hooks for several actions that cot performs:

	link
	install
	update
	unlink
	uninstall

Each of these actions can be hooked at three times: before, during, and after.
Before and after are prefixed with 'pre-' and 'post-', respectively.
During is named as is. For, example, files checked when linking are:

	pre-link
	link
	post-link

The file will only be executed if it has precisely this name, has the
executable bit set, and can be executed by the kernel.
`))

var ErrRepoExists = errors.New("cot repository already exists")

// Init initializes a new cot repository in the directory specified by path.
// If path is within an existing cot repository, then an error RepoExistsErr is returned.
func Init(path string) error {
	// Don't initialize if within existing repository
	if found, err := FindRoot(path); err != nil || found != "" {
		if found != "" {
			return ErrRepoExists
		}
		return err
	}
	abspath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Create configuration directory
	writeTmpl := func(path string, tmpl *template.Template, obj interface{}) error {
		path = filepath.Join(CotDir[0], path)
		dir, _ := filepath.Split(path)
		ex, err := osutil.DirExists(dir)
		if err != nil {
			return err
		}
		if !ex {
			log.Debugln("MkdirAll", dir)
			os.MkdirAll(dir, 0755)
		}
		log.Debugln("Create", path)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		return tmpl.Execute(f, obj)
	}

	name := filepath.Base(abspath)
	writeTmpl(CotDirConfig[0], configTmpl, struct{ Name string }{Name: name})
	writeTmpl(CotDirExclude[0], excludeTmpl, nil)
	writeTmpl(CotDirDirectory[0], directoryTmpl, nil)
	writeTmpl(filepath.Join(CotDirHooks[0], "README"), hooksReadMeTmpl, nil)
	log.Infoln("Initialized empty cot repository", filepath.Join(abspath, CotDir[0]))
	return nil
}

// IsRoot checks if the given path represents the root of a cot repository.
// If any error is returned during filesystem operations, then an error
// is returned.
func IsRoot(path string) (bool, error) {
	// Check if any valid configuration directories exist:
	for _, cd := range CotDir {
		ex, err := osutil.DirExists(filepath.Join(path, cd))
		if err != nil {
			return false, err
		}
		if ex {
			return true, nil
			// TODO: check if the directory contains a properly initialized
			// cot repository.
		}
	}

	// Check if any valid configuration files exist:
	for _, cf := range CotConfig {
		ex, err := osutil.FileExists(filepath.Join(path, cf))
		if err != nil {
			return false, err
		}
		if ex {
			return true, nil
			// TODO: check if the file contains a valid configuration.
		}
	}

	return false, nil
}

// FindRoot finds the path to the root of the cot containing path,
// or "" if path is not within a cot.
func FindRoot(path string) (string, error) {
	return pathutil.FindRootDir(path, IsRoot)
}
