/**
*
* @file: install.go
* @description: This file contains the install command to run the necessary install scripts.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package cmd

import (
	"github.com/spf13/cobra"
	"tomavilius.in/atlas/internal/atlasinstall"
)

func init() {

	rootCommand.AddCommand(installCommand)
}

// install command to run install scripts.
var installCommand = &cobra.Command{

	Use:   "install",
	Short: "Run installation procedure for atlas.",
	Long:  "Run the necessary install scripts. Don't run again and again.",
	Run: func(cmd *cobra.Command, args []string) {

		// transfer control to the function to run the install scripts.
		atlasinstall.InstallAtlas()
	},
}
