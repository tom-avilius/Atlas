/**
*
* @file: watch.go
* @description: This file contains the watch commands to attach fs notify that checks for changes and takes necessary action.
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

	rootCommand.AddCommand(watchCommand)
}

// watchCommand watches for changes in the filesystem to take necessary action.
var watchCommand = &cobra.Command{

	Use:   "watch",
	Short: "Starts to watch all the local folder and sub-folders within.",
	Long:  "If any change occurs within any local folder then it performs git add .",
	Run: func(cmd *cobra.Command, args []string) {

		// transfer to AttachFsNotify.
		reposync.AttachFsNotify()
	},
}
