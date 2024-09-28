package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


func init () {

  rootCommand.AddCommand(addCommand)
}

var addCommand = &cobra.Command {
  
  Use: "add",
  Short: "Add a git repository",
  Long: "Add a git repository for atlas to manage and make backups to.",
  Run: func(cmd *cobra.Command, args []string) {

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Name: ")
    repoName, _ := reader.ReadString('\n');

    fmt.Print("Url: ")
    repoUrl, _ := reader.ReadString('\n')

    repoName = repoName[:len(repoName)-1]
    repoUrl = repoUrl[:len(repoUrl)-1]

    fmt.Println("\nConfirm Action:")
    fmt.Println("Name:\t" +repoName + "\t(need not be the github repo name)")
    fmt.Println("Url:\t" +repoUrl + "\t(needs to be the github clone link)")
    
    fmt.Println("\nConfirm Action by pressing y or n to abort.")
    choice, _ := reader.ReadString('\n');

    if len(choice) > 2 {

      fmt.Println("Too long an input.")
      os.Exit(1)
    }

    if choice[0] == 'y' || choice[0] == 'Y' {

      fmt.Println("\nProceeding to clone the repository..")
    } else if choice[0] == 'n' || choice[0] == 'N' {

      fmt.Println("\nAborting.. \nNo changes have been made.")
    } else {

      fmt.Println("\nNo option selected. \nAborting..")
    }
  },
}
