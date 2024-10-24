package cmd

import (
	"github.com/spf13/cobra"
	"tomavilius.in/atlas/internal/reposync"
)



func init () {

  rootCommand.AddCommand(syncCommand)
}


var syncCommand = &cobra.Command {

  Use: "sync",
  Short: "Sync the local folders with github repository.",
  Long: "Performs a git push if any local changes arise for all the repositories.",
  Run: func(cmd *cobra.Command, args []string) {
  
    // writing path data
    // TODO: Inside the function have some method to check for 
    // changes and only then write to the file.
    reposync.WritePathData()
  },
}
