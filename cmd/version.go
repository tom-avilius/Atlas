/**
*
* @file: version.go
* @description: This file contains the version command to display the atlas version.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

	rootCommand.AddCommand(versionCommand)
}

// versionCommand to display the version of atlas.
var versionCommand = &cobra.Command{

	Use:   "version",
	Short: "Displays the version.",
	Long:  "Displays the installed version of the atlas project.",
	Run: func(cmd *cobra.Command, args []string) {

		// print the current version
		// HACK: Ain't there a better way to do this shit, than just hardcoding it?
		fmt.Println("v0.0.1 - development")
	},
}
