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
  Run: func(cmd *cobra.Command, args []string) {
    
    fmt.Println("Usage atlas [command]")

    fmt.Println("\nAvailable Commands:")
    fmt.Println("version \t Display the atlas project version.")
    fmt.Println("add \t Add a git repository for atlas to make backups to.")
    fmt.Println("remove \t Remove a git repository for atlas to stop making backups to.")
    fmt.Println("sync \t Explicitly sync all the repositories managed by atlas.")
    fmt.Println("watch \t Re-evaluate path file and attach fsnotify to them.")
    fmt.Println("list \t Lists all the repositories managed by atlas.")
    fmt.Println("info \t Detailed list view of atlas config.")
  },
}

func Execute () {

  err := rootCommand.Execute()

  if err != nil {

    fmt.Println(err)
    os.Exit(1)
  }
}
