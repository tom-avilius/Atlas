
package cmd

import (

  "github.com/spf13/cobra"
  "tomavilius.in/atlas/internal/reposync"
)



func init () {

  rootCommand.AddCommand(watchCommand)
}


var watchCommand = &cobra.Command {

  Use: "watch",
  Short: "Starts to watch all the local folder and sub-folders within.",
  Long: "If any change occurs within any local folder then it performs git add .",
  Run: func(cmd *cobra.Command, args []string) {

    reposync.AttachFsNotify()
  },
}

