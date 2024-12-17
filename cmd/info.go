/**
*
* @file: info.go
* @description: This file contains the add command that adds a repository for atlas to watch and make backups to.
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

// add info command to the root command.
func init() {

	rootCommand.AddCommand(infoCommand)
}

// info command to give detailed view of atlas config. Prints everything.
var infoCommand = &cobra.Command{

	Use:   "info",
	Short: "Detailed list view of atlas config.",
	Long:  "Lists in detail the config file of atlas.",
	Run: func(cmd *cobra.Command, args []string) {

		// transfer control to DetailedRepositoryView function to print entire atlas config.
		repoinformer.DetailedRepositoryView()
	},
}
