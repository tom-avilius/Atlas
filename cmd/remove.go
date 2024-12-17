/**
*
* @file: add.go
* @description: This file contains the remove command to remove a repository from config. Opposite to add command.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"tomavilius.in/atlas/internal/reporegistry"

	"github.com/spf13/cobra"
)

func init() {

	rootCommand.AddCommand(removeCommand)

	// --clear | -c | To clear the entire config file, essentially just replaces it with a new empty one.
	// defaults to false
	removeCommand.Flags().BoolP("clear", "c", false, "Clear the entire config.")
}

// the remove command to remove repo from config.
var removeCommand = &cobra.Command{

	Use:   "remove",
	Short: "Removes a repository locally.",
	Long:  "Remove a git repository or stop managing it. Only works locally. Would not delete the remote repository.",
	Run: func(cmd *cobra.Command, args []string) {

		// checking -c | --clear flag
		if clearConfig, _ := cmd.Flags().GetBool("clear"); clearConfig {

			// log the action to clear the entire config.
			fmt.Println("Clearing the entire config.")
			fmt.Println("Note: This does not affect any cloned folders, or anything other than the config in general.")

			// clear the config.
			reporegistry.ClearConfig()

			os.Exit(0)
		}

		// when no flag was passed:

		// create the reader.
		reader := bufio.NewReader(os.Stdin)

		// input the repository name to remove.
		fmt.Print("Repository Name: ")
		repoName, _ := reader.ReadString('\n')
		repoName = strings.ToLower(repoName)

		// invoke DeleteRepository to actually do the deleting.
		reporegistry.DeleteRepository(true, strings.TrimSpace(repoName))
	},
}
