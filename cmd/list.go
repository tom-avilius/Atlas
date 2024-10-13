
package cmd


import (
	"github.com/spf13/cobra"
	"tomavilius.in/atlas/internal/repoinformer"
)



func init () {

  rootCommand.AddCommand(listCommand)
}


var listCommand = &cobra.Command {

  Use: "list",
  Short: "List only the repositories.",
  Long: "Lists all the repositories managed by atlas.",
  Run: func(cmd *cobra.Command, args []string) {

    repoinformer.ListRepositories()
  },
}

