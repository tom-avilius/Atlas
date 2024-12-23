/**
*
* @file: root.go
* @description: The entry point for cobra cmd, the root command -> atlas.
* 							prints out all the commands and their short description.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

// the cmd package that contains all the commands available.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// creating the root command to print all the available commands.
var rootCommand = &cobra.Command{

	Use:   "atlas",
	Short: "A backup tool for you configs and notes.",
	Long: `Atlas uses git to backup the files and folders.
You must first create a git repository and provide the url for atlas to do its thing.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Usage atlas [command]")

		fmt.Println("\nAvailable Commands:")
		fmt.Println("version \t Display the atlas project version.")
		fmt.Println("add \t Add a git repository for atlas to make backups to.")
		fmt.Println("remove \t Remove a git repository for atlas to stop making backups to.")
		fmt.Println("sync \t Explicitly sync all the repositories managed by atlas.")
		fmt.Println("watch \t Re-evaluate path file and attach fsnotify to them.")
		fmt.Println("list \t Lists all the repositories managed by atlas.")
		fmt.Println("info \t Detailed list view of atlas config.")
		fmt.Println("install \t Run the necessary install scripts.")
		fmt.Println("start \t Start atlas service along with necessary components.")
	},
}

// entry point.
func Execute() {

	err := rootCommand.Execute()

	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}
}
