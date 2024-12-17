/**
*
* @file: add.go
* @description: This file contains the add command that adds a repository for atlas to watch and make backups to.
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
	"path"
	"strings"
	"time"

	"tomavilius.in/atlas/internal/reporegistry"

	"github.com/spf13/cobra"
)

// adds the add command to root command.
func init() {

	rootCommand.AddCommand(addCommand)
}

// creating the add command
var addCommand = &cobra.Command{

	Use:   "add",
	Short: "Add a git repository",
	Long:  "Add a git repository for atlas to manage and make backups to.",
	Run: func(cmd *cobra.Command, args []string) {

		// start with the logs to explain the command behaviour.
		fmt.Println("\nThis command will walk you over adding a new git repository to atlas.")
		fmt.Println("The specified local folder or file would be backed up to this repository.")

		// initialize the reader to take input from the console.
		reader := bufio.NewReader(os.Stdin)

		// take in the name of the repository.
		// NOTE: This is for local use only.
		fmt.Print("\nName: ")
		repoName, _ := reader.ReadString('\n')
		repoName = strings.ToLower(repoName)

		// take in the remote url of the repository.
		// TEST: Would it work if the clone link is provided?
		fmt.Print("Url: ")
		repoUrl, _ := reader.ReadString('\n')

		// to store the path to clone the repository to.
		var clonePath string
		// getting the home directory.
		// PERF: Use go's way of using if statements.
		homeDir, err := os.UserHomeDir()
		if err == nil {

			// display the home directory and then read the rest from the console.
			fmt.Print("Clone Path: " + homeDir)
			clonePath, _ = reader.ReadString('\n')
		} else {

			// TODO: This logging might be redundant. Check it.
			fmt.Println("Error occured while getting home directory.")
			fmt.Println("Error Log:")
			fmt.Println(err)
			os.Exit(1)
		}

		// trimming the strings.
		repoName = repoName[:len(repoName)-1]
		repoUrl = repoUrl[:len(repoUrl)-1]
		clonePath = clonePath[:len(clonePath)-1]
		// updating clone path with the home dir.
		clonePath = homeDir + clonePath

		// asking to confirm the action to continue to clone the repository.
		fmt.Println("\nConfirm Action:")
		fmt.Println("Name:\t" + repoName + "\t(need not be the github repo name)")
		fmt.Println("Url:\t" + repoUrl + "\t(needs to be the github repo link)")
		fmt.Println("Path:\t" + clonePath + "\t(the local clone path)")

		// reader user choice to confirm or abort.
		fmt.Println("\nConfirm Action by pressing y or n to abort.")
		choice, _ := reader.ReadString('\n')

		// if invalid choice is made.
		// PERF: Use go's way of writing if statements. Nest with the next one in this case.
		if len(choice) > 2 {

			fmt.Println("Too long an input. \nAborting..")
			os.Exit(1)
		}

		// if the user proceeds then create a Repository with all the data.
		if choice[0] == 'y' || choice[0] == 'Y' {

			repo := reporegistry.Repository{

				Name:     strings.TrimSpace(repoName),
				Url:      repoUrl,
				Path:     path.Join(clonePath, repoName),
				AddedAt:  time.Now(),
				LastSync: time.Time{},
			}

			// add the repository. Control is transfered to the function handling everything.
			reporegistry.AddRepository(repo, clonePath)

			// if the user chose to abort.
		} else if choice[0] == 'n' || choice[0] == 'N' {

			fmt.Println("\nAborting.. \nNo changes have been made.")
		} else {

			// when only enter was pressed, just abort.
			fmt.Println("\nNo option selected. \nAborting..")
		}
	},
}
