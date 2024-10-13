package cmd

import (
	"github.com/spf13/cobra"
	"tomavilius.in/atlas/internal/repoinformer"
)


func init () {

  rootCommand.AddCommand(infoCommand)
}

var infoCommand =  &cobra.Command {

  Use: "info",
  Short: "Detailed list view of atlas config.",
  Long: "Lists in detail the config file of atlas.",
  Run: func(cmd *cobra.Command, args []string) {

    repoinformer.DetailedRepositoryView()
  },
}
