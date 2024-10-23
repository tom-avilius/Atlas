package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"tomavilius.in/atlas/internal/reporegistry"

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

    fmt.Println("\nThis command will walk you over adding a new git repository to atlas.")
    fmt.Println("The specified local folder or file would be backed up to this repository.")

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("\nName: ")
    repoName, _ := reader.ReadString('\n');
    repoName = strings.ToLower(repoName)

    fmt.Print("Url: ")
    repoUrl, _ := reader.ReadString('\n')

    var clonePath string
    homeDir, err := os.UserHomeDir()
    if err == nil {

      fmt.Print("Clone Path: " +homeDir);
      clonePath, _ = reader.ReadString('\n')
    } else {

      fmt.Println("Error occured while getting home directory.")
      fmt.Println("Error Log:")
      fmt.Println(err);
      os.Exit(1)
    }

    repoName = repoName[:len(repoName)-1]
    repoUrl = repoUrl[:len(repoUrl)-1]
    clonePath = clonePath[:len(clonePath)-1]
    clonePath = homeDir+clonePath

    fmt.Println("\nConfirm Action:")
    fmt.Println("Name:\t" +repoName + "\t(need not be the github repo name)")
    fmt.Println("Url:\t" +repoUrl + "\t(needs to be the github repo link)")
    fmt.Println("Path:\t" +clonePath +"\t(the local clone path)")
    
    fmt.Println("\nConfirm Action by pressing y or n to abort.")
    choice, _ := reader.ReadString('\n');

    if len(choice) > 2 {

      fmt.Println("Too long an input. \nAborting..")
      os.Exit(1)
    }

    if choice[0] == 'y' || choice[0] == 'Y' {
      
      repo := reporegistry.Repository {
        
        Name: strings.TrimSpace(repoName),
        Url: repoUrl,
        AddedAt: time.Now(),
        LastSync: time.Time{},
      } 

      reporegistry.AddRepository(repo, clonePath)
    } else if choice[0] == 'n' || choice[0] == 'N' {

      fmt.Println("\nAborting.. \nNo changes have been made.")
    } else {

      fmt.Println("\nNo option selected. \nAborting..")
    }
  },
}
