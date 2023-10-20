/*
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
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/portworx/pxc/pkg/commander"
	pluginpkg "github.com/portworx/pxc/pkg/plugin"
	"github.com/portworx/pxc/pkg/util"

	"github.com/spf13/cobra"
)

type pluginListOptions struct {
	Verifier PathVerifier
	NameOnly bool
	Lister   pluginpkg.PluginLister

	PluginPaths []string
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

var (
	pluginListLong = `
		List all available component files on a user's PATH.

		Available component files are those that are:
		- executable
		- anywhere on the user's PATH
		- begin with "pxc-"`

	listOptions *pluginListOptions
	listCmd     *cobra.Command
)

var _ = commander.RegisterCommandVar(func() {
	listOptions = &pluginListOptions{}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all visible component executables",
		Long:  pluginListLong,
		RunE:  listExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	listCmd.Flags().BoolVar(&listOptions.NameOnly, "name-only", listOptions.NameOnly, "If true, display only the binary name of each component, rather than its full path")
	PluginAddCommand(listCmd)
})

func ListAddCommand(cmd *cobra.Command) {
	listCmd.AddCommand(cmd)
}

func listExec(cmd *cobra.Command, args []string) error {
	listOptions.Lister.Complete(cmd)
	return listOptions.Run(cmd, args)
}

func (o *pluginListOptions) Run(cmd *cobra.Command, args []string) error {
	isFirstFile := true

	components, err := o.Lister.GetList()
	if err != nil {
		return err
	}

	if len(components) == 0 {
		return fmt.Errorf("Unable to find any pxc components in your PATH")
	}

	for _, component := range components {
		if isFirstFile {
			util.Eprintf("The following compatible components are available:\n\n")
			isFirstFile = false
		}
		util.Printf("%s\n", component.Path)
		for _, warning := range component.Warnings {
			util.Eprintf("  - %s\n", warning)
		}
	}

	return nil
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
