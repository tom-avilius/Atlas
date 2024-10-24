package reposync

import (
	"fmt"
	"os"
	"path/filepath"

	// "github.com/fsnotify/fsnotify"
	"tomavilius.in/atlas/internal/reporegistry"
)



func dirList (path string) ([]string, bool) {

  if filePath, success := reporegistry.HandleHomeDirectory(path); success {

    path = filePath
  } else {

    return nil, false
  }

  if _, err := os.Stat(path); err != nil {

    fmt.Println("Error while pinging directory.")
    fmt.Println("Was about to check for child directories.")

    if (os.IsNotExist(err)) {

      fmt.Println("The directory does not exist.")
    }
    fmt.Println("Error Log:")
    fmt.Println(err)

    return nil, false
  }

  var directories []string

  error := filepath.Walk(path, func(dir string, info os.FileInfo, err error) error {

    if err != nil {

      return err
    }

    if info.IsDir() && path != dir {
      
      directories = append(directories, dir)
    }

    return nil
  })

  if error != nil {

    fmt.Println("Error occured while getting children directories.")
    fmt.Println("Error Log:")
    fmt.Println(error)
    return nil, false
  }

  return directories, true
}
