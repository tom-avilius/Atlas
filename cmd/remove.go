package cmd

import (
	"bufio"
	"fmt"
	"os"
  "tomavilius.in/atlas/internal/reporegistry"

	"github.com/spf13/cobra"
)


func init () {

  rootCommand.AddCommand(removeCommand)
}


var removeCommand = &cobra.Command {

  Use: "remove",
  Short: "Removes a repository locally.",
  Long: "Remove a git repository or stop managing it. Only works locally. Would not delete the remote repository.",
  Run: func(cmd *cobra.Command, args []string) {

    fmt.Print("Repository Name: ");
    
    reader := bufio.NewReader(os.Stdin)

    repoName, _ := reader.ReadString('\n')

    // TODO: match the reponame with all the names in the config file, if the repo is not there inform the user.

    reporegistry.DeleteRepository(true, repoName)
  },
}
