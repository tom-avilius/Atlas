package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


func init () {

  rootCommand.AddCommand(installCommand)
}


var installCommand = &cobra.Command {

  Use: "install",
  Short: "Run installation procedure for atlas.",
  Long: "Run the necessary install scripts. Don't run again and again.",
  Run: func(cmd *cobra.Command, args []string) {

    
  },
}
