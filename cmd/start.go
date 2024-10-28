package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


func init () {

  rootCommand.AddCommand(startCommand)
}


var startCommand = &cobra.Command {

  Use: "start",
  Short: "Start the atlas daemon.",
  Long: "Start atlas service along with necessary components.",
  Run: func(cmd *cobra.Command, args []string) {

    fmt.Println("hehe")
  },
}
