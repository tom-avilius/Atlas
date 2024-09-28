package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


func init () {

  rootCommand.AddCommand(versionCommand)
}


var versionCommand = &cobra.Command{

  Use: "version",
  Short: "Displays the version.",
  Long: "Displays the installed version of the atlas project.",
  Run: func (cmd *cobra.Command, args []string) {
    
    fmt.Println("v0.0.1 - development")
  },
}

