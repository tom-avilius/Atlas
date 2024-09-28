package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var rootCommand = &cobra.Command{

  Use: "atlas",
  Short: "A backup tool for you configs and notes.",
  Long: `Atlas uses git to backup the files and folders.
You must first create a git repository and provide the url for atlas to do its thing.`,
}

func Execute () {

  err := rootCommand.Execute()

  if err != nil {

    fmt.Println(err)
    os.Exit(1)
  }
}
