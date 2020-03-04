/*
Copyright © 2019 Portworx

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
package component

import (
	"fmt"
	"os"

	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

type gendocCmdArgs struct {
	outputDir string
	format    string
}

var (
	gendocCmd  *cobra.Command
	gendocArgs *gendocCmdArgs
)

var _ = commander.RegisterCommandVar(func() {
	gendocArgs = &gendocCmdArgs{}
	gendocCmd = &cobra.Command{
		Use:     "gendoc",
		Aliases: []string{"gendocs"},
		Short:   "Generate doc files in Markdown format",
		// Hide this command. Only used for generating docs by developers
		Hidden: true,
		RunE:   gendocExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	RootAddCommand(gendocCmd)

	gendocCmd.Flags().StringVar(&gendocArgs.format, "format", "md", "Output formats: md, man")
	gendocCmd.Flags().StringVar(&gendocArgs.outputDir, "output-dir", "pxdocs", "Output directory")
})

func GenDocAddCommand(cmd *cobra.Command) {
	gendocCmd.AddCommand(cmd)
}

func gendocExec(cmd *cobra.Command, args []string) error {
	switch gendocArgs.format {
	case "md":
		util.Printf("Creating markdown docs in %s...\n", gendocArgs.outputDir)

		os.MkdirAll(gendocArgs.outputDir, 0755)
		return doc.GenMarkdownTree(rootCmd, gendocArgs.outputDir)
	case "man":
		util.Printf("Creating man pages in %s...\n", gendocArgs.outputDir)

		os.MkdirAll(gendocArgs.outputDir, 0755)
		manHeader := &doc.GenManHeader{
			Title: "pxc Portworx client",
		}
		return doc.GenManTree(rootCmd, manHeader, gendocArgs.outputDir)

	default:
		return fmt.Errorf("Unknown format: %s", gendocArgs.format)
	}

}
