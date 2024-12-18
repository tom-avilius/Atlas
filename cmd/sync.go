/**
*
* @file: sync.go
* @description: This file contains the sync command to push to git repository if any change arises.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package cmd

import (
	"github.com/spf13/cobra"
	"tomavilius.in/atlas/internal/reposync"
)

func init() {

	rootCommand.AddCommand(syncCommand)
}

// TODO: Implement the push and commit function in this command.
// TEST: Does the add functionality even work?

// sync command to push local changes to the remote repository.
var syncCommand = &cobra.Command{

	Use:   "sync",
	Short: "Sync the local folders with github repository.",
	Long:  "Performs a git push if any local changes arise for all the repositories.",
	Run: func(cmd *cobra.Command, args []string) {

		// writing path data
		// TODO: Inside the function have some method to check for
		// changes and only then write to the file.
		// TEST: I guess the above is implemented but ensure that it works.
		reposync.WritePathData()
	},
}
