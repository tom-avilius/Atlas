package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"tomavilius.in/atlas/internal/reporegistry"

	"github.com/spf13/cobra"
)


func init () {

  rootCommand.AddCommand(removeCommand)

  // --clear | -c | To clear the entire config file, essentially just replaces it with a new empty one.
  // defaults to false
  removeCommand.Flags().BoolP("clear", "c", false, "Clear the entire config.")
}


var removeCommand = &cobra.Command {

  Use: "remove",
  Short: "Removes a repository locally.",
  Long: "Remove a git repository or stop managing it. Only works locally. Would not delete the remote repository.",
  Run: func(cmd *cobra.Command, args []string) {

    // checking -c | --clear flag
    if clearConfig, _ := cmd.Flags().GetBool("clear"); clearConfig {

      fmt.Println("Clearing the entire config.")
      fmt.Println("Note: This does not affect any cloned folders, or anything other than the config in general.")
      os.Exit(0)
    }

    fmt.Print("Repository Name: ");
    
    reader := bufio.NewReader(os.Stdin)

    repoName, _ := reader.ReadString('\n')

    repoName = strings.ToLower(repoName)

    reporegistry.DeleteRepository(true, strings.TrimSpace(repoName))
  },
}
