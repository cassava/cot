package cot

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

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

const shortConfigTmpl = template.Must(template.New("cot-short-config").Parse(
	`# cot repository configuration
name = {{.Name}}
exclude = ['.*']
directory = ['.local', '.local/*', '.config'
`))

const configTmpl = template.Must(template.New("cot-config").Parse(
	`# cot repository configuration
name = {{.Name}}
`))

const excludeTmpl = template.Must(template.New("cot-exclude").Parse(
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

var RepoExistsErr = errors.New("cot repository already exists")

// Init will only initialize in a
func Init(path, name string) error {
	if found, err := FindRoot(path); err != nil || found {
		if found {
			return RepoExistsErr
		}
		return err
	}

	// Create cot configuration directory
	writeTmpl = func(path, tmpl *template.Template, obj interface{}) error {
		path = filepath.Join(CotDir[0], path)
		dir, file = filepath.Split(path)
		ex, err := osutil.DirExists(dir)
		if err != nil {
			return err
		}
		if !ex {
			os.MkdirAll(dir)
		}
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		return tmpl.Execute(f, obj)
	}

	writeTmpl(CotDirConfig[0], configTmpl, struct{ Name string }{Name: name})
	writeTmpl(CotDirExclude[0], excludeTmpl, nil)
	writeTmpl(CotDirDirectory[0], directoryTmpl, nil)
	writeTmpl(filepath.Join(CotDirHooks[0], hooksReadMeTmpl, nil))
}

func InitShort(path, name string) error {
	if found, err := FindRoot(path); err != nil || found {
		if found {
			return RepoExistsErr
		}
		return err
	}

	// TODO: Create short cot configuration file

}

// IsRoot checks if the given path represents the root of a cot repository.
// If any error is returned during filesystem operations, then an error
// is returned.
//
// TODO: There are multiple file names that we can check.
func IsRoot(path string) (bool, error) {
	cotdir := filepath.Join(path, CotDir)
	ex, err := osutil.DirExists(cotdir)
	if err != nil {
		return false, err
	}
	// TODO: check if the directory contains a properly initialized
	// cot repository.
	return ex, nil
}

// FindRoot finds the first cot directory in the directory hierarchy.
// This is determined by the existence of either a .cot directory or
// a .cotconfig file, which is supposed to be in the root of a dungeon,
// just like it is with Git.
//
// TODO: This hasn't been modified for cot yet!
func FindRoot(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	for {
		found, err := IsRoot(path)
		if err != nil {
			return "", err
		} else if found {
			return path, nil
		}

		// After we have checked / we stop looking.
		if path == "/" {
			break
		}
		path = filepath.Dir(path)
	}

	return "", fmt.Errorf("could not find dungeon root in %q", path)
}
