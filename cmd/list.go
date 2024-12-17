/**
*
* @file: list.go
* @description: This file contains the list command to list only the repo names from the config file.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package cmd

import (
	"github.com/spf13/cobra"
	"tomavilius.in/atlas/internal/repoinformer"
)

func init() {

	rootCommand.AddCommand(listCommand)
}

// list command to list only repo names from atlas config.
var listCommand = &cobra.Command{

	Use:   "list",
	Short: "List only the repositories.",
	Long:  "Lists all the repositories managed by atlas.",
	Run: func(cmd *cobra.Command, args []string) {

		// transfer control to function to list repositories only.
		repoinformer.ListRepositories()
	},
}
