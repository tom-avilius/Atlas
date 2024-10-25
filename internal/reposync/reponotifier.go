
package reposync

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"tomavilius.in/atlas/internal/reporegistry"
)



// returns a list of child directories
func dirList (path string) ([]string, bool) {

  // handling ~ paths
  if filePath, success := reporegistry.HandleHomeDirectory(path); success {

    path = filePath
  } else {

    return nil, false
  }

  // checking if the path exists
  if _, err := os.Stat(path); err != nil {

    // if an error occurs.
    fmt.Println("Error while pinging directory.")
    fmt.Println("Was about to check for child directories.")

    // if path also does not exist.
    if (os.IsNotExist(err)) {

      fmt.Println("The directory does not exist.")
    }
    fmt.Println("Error Log:")
    fmt.Println(err)

    return nil, false
  }

  // to store the list of directories
  var directories []string

  // traversing through the entire directory tree
  error := filepath.Walk(path, func(dir string, info os.FileInfo, err error) error {

    if err != nil {

      return err
    }

    // checking if the entity is a directory
    if info.IsDir() && path != dir {
      
      // appending it to the list of directories
      directories = append(directories, dir)
    }

    // successful function return
    return nil
  })

  // if an error occurs while traversing the directory tree
  if error != nil {

    fmt.Println("Error occured while getting children directories.")
    fmt.Println("Error Log:")
    fmt.Println(error)
    return nil, false
  }

  // successful function ends successfulLy
  return directories, true
}


// function to attach fsnotify to the dirlist
func attachFsnotify (dirList []string) {

  watcher, err := fsnotify.NewWatcher()
  if err != nil {

    fmt.Println("Error occured while creating watcher.")
    fmt.Println("Error Log:")
    fmt.Println(err);
    os.Exit(1)
  }
  defer watcher.Close()

  go func () {

    for {

      select {

      case event, ok := <-watcher.Events:
        if !ok {

          return
        }
        fmt.Println("event:", event)
        if event.Has(fsnotify.Write) {

          fmt.Println("Modified File");
          fmt.Println("File Path: " +event.Name);
        }
        if event.Has(fsnotify.Remove) {

          fmt.Println("File Removed");
          fmt.Println("File Path: " +event.Name)
        }
        if event.Has(fsnotify.Create) {

          fmt.Println("New Path Created")
          fmt.Println("File Path: " +event.Name)
        }
      case err, ok := <-watcher.Errors:
        if !ok {

          return
        }
        fmt.Println("Error: ")
        fmt.Println(err)
      }
    }
  } ()

  for _, dir := range dirList {
    
    err = watcher.Add(dir)
    if err != nil {

      fmt.Println(err)
      os.Exit(1)
    }
  }

  <-make(chan struct{})
}
