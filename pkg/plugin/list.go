/*
Copyright 2021 Pure Storage
Copyright 2017 The Kubernetes Authors.

Originally from:
https://raw.githubusercontent.com/kubernetes/pxc/release-1.17/pkg/cmd/plugin/plugin.go

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package plugin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"

	"github.com/spf13/cobra"
)

type PluginLister struct {
	Verifier PathVerifier
	NameOnly bool

	PluginPaths []string
}

type Component struct {
	Path     string
	Warnings []string
}

// PathVerifier receives a path and determines if it is valid or not
type PathVerifier interface {
	// Verify determines if a given path is valid
	Verify(path string) []error
}

type CommandOverrideVerifier struct {
	root        *cobra.Command
	seenPlugins map[string]string
}

func (o *PluginLister) Complete(cmd *cobra.Command) error {
	o.Verifier = &CommandOverrideVerifier{
		root:        cmd.Root(),
		seenPlugins: make(map[string]string),
	}

	o.PluginPaths = filepath.SplitList(os.Getenv("PATH"))
	o.PluginPaths = append(o.PluginPaths, path.Join(config.CM().GetFlags().ConfigDir, "bin"))
	return nil
}

func (o *PluginLister) GetList() ([]*Component, error) {
	pluginErrors := []error{}
	components := make([]*Component, 0)

	for _, dir := range uniquePathsList(o.PluginPaths) {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			if _, ok := err.(*os.PathError); ok {
				continue
			}

			pluginErrors = append(pluginErrors, fmt.Errorf("error: unable to read directory %q in your PATH: %v", dir, err))
			continue
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if !hasValidPrefix(f.Name(), ValidPluginFilenamePrefixes) {
				continue
			}

			pluginPath := f.Name()
			if !o.NameOnly {
				pluginPath = filepath.Join(dir, pluginPath)
			}

			component := &Component{
				Path:     pluginPath,
				Warnings: make([]string, 0),
			}

			if errs := o.Verifier.Verify(filepath.Join(dir, f.Name())); len(errs) != 0 {
				for _, err := range errs {
					component.Warnings = append(component.Warnings, err.Error())
				}
			}
			components = append(components, component)
		}
	}

	if len(pluginErrors) > 0 {
		util.Eprintf("\n")
		errs := bytes.NewBuffer(nil)
		for _, e := range pluginErrors {
			fmt.Fprintln(errs, e)
		}
		return nil, fmt.Errorf("%s", errs.String())
	}

	return components, nil
}

func (o *PluginLister) GetSortedRootComponents() ([]string, error) {
	components, err := o.GetList()
	if err != nil {
		return nil, err
	}

	list := make([]string, 0)

	// Add components as subcommands
	for _, component := range components {

		// Remove the "pxc-" part of the component
		c := strings.TrimPrefix(component.Path, "pxc-")

		// Only add top level component, no sub components
		if strings.Contains(c, "-") {
			continue
		}

		list = append(list, c)
	}

	sort.Strings(list)
	return list, nil
}

// Verify implements PathVerifier and determines if a given path
// is valid depending on whether or not it overwrites an existing
// pxc command path, or a previously seen plugin.
func (v *CommandOverrideVerifier) Verify(path string) []error {
	if v.root == nil {
		return []error{fmt.Errorf("unable to verify path with nil root")}
	}

	// extract the plugin binary name
	segs := strings.Split(path, "/")
	binName := segs[len(segs)-1]

	cmdPath := strings.Split(binName, "-")
	if len(cmdPath) > 1 {
		// the first argument is always "pxc" for a plugin binary
		cmdPath = cmdPath[1:]
	}

	errors := []error{}

	if isExec, err := isExecutable(path); err == nil && !isExec {
		errors = append(errors, fmt.Errorf("warning: %s identified as a pxc component, but it is not executable", path))
	} else if err != nil {
		errors = append(errors, fmt.Errorf("error: unable to identify %s as an executable file: %v", path, err))
	}

	if existingPath, ok := v.seenPlugins[binName]; ok {
		errors = append(errors, fmt.Errorf("warning: %s is overshadowed by a similarly named component: %s", path, existingPath))
	} else {
		v.seenPlugins[binName] = path
	}

	if cmd, _, err := v.root.Find(cmdPath); err == nil {
		errors = append(errors, fmt.Errorf("warning: %s overwrites existing command: %q", binName, cmd.CommandPath()))
	}

	return errors
}

func isExecutable(fullPath string) (bool, error) {
	info, err := os.Stat(fullPath)
	if err != nil {
		return false, err
	}

	if runtime.GOOS == "windows" {
		fileExt := strings.ToLower(filepath.Ext(fullPath))

		switch fileExt {
		case ".bat", ".cmd", ".com", ".exe", ".ps1":
			return true, nil
		}
		return false, nil
	}

	if m := info.Mode(); !m.IsDir() && m&0111 != 0 {
		return true, nil
	}

	return false, nil
}

// uniquePathsList deduplicates a given slice of strings without
// sorting or otherwise altering its order in any way.
func uniquePathsList(paths []string) []string {
	seen := map[string]bool{}
	newPaths := []string{}
	for _, p := range paths {
		if seen[p] {
			continue
		}
		seen[p] = true
		newPaths = append(newPaths, p)
	}
	return newPaths
}

func hasValidPrefix(filepath string, validPrefixes []string) bool {
	for _, prefix := range validPrefixes {
		if !strings.HasPrefix(filepath, prefix+"-") {
			continue
		}
		return true
	}
	return false
}
