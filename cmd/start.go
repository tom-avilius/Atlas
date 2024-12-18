/**
*
* @file: start.go
* @description: This file contains the start command to start the necessary start up components for the atlas daemon.
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

	rootCommand.AddCommand(startCommand)
}

// TODO: Implement the start command.

// start command to start the necessary atlas components
var startCommand = &cobra.Command{

	Use:   "start",
	Short: "Start the atlas daemon.",
	Long:  "Start atlas service along with necessary components.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("hehe")
	},
}
