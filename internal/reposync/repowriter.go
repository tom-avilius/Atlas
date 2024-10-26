package reposync

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
	"tomavilius.in/atlas/internal/reporegistry"
)

// function to create a path file which stores the various paths that must be added to fsnotify
func createPathFile() bool {

  // if the path file does not exist, create it
  if !reporegistry.CheckFileExist(reporegistry.PathFilePath) {

    fmt.Println("\nPath file does not exist.")
    fmt.Println("Creating path file..")
    // if atlas config directory could not be created
    if success := reporegistry.CreateDir(reporegistry.ConfigDir); !success {

      fmt.Println("Could not create atlas config directory.")
      return false
    }
    
    // if path file could not be created
    if success := reporegistry.CreateFile(reporegistry.PathFilePath); !success {

      fmt.Println("Could not create path file for atlas.")
      return false
    }

    fmt.Println("Created path file.")
  }

  // successful
  return true
} 


// function to write path data to the path file.
// returns false if an error occurs.
func writePathData (filepath string, pathData reporegistry.PathData) bool {

  // handling ~ paths
  if path, success := reporegistry.HandleHomeDirectory(filepath); success {

    filepath = path
  } else {
    
    // unsuccessful 
    return false
  }

  // to store yaml data.
  var data reporegistry.PathData

  // read the yaml file
  content, err := os.ReadFile(filepath)
  // Log if an error occurs
  if err != nil {

    fmt.Println("Error occured while reading path file.")
    fmt.Println(err)
    return false
  }

  // Unmarshal the yaml data
  err = yaml.Unmarshal(content, &data)
  // Log if an error occurs
  if err != nil {

    fmt.Println("Error occured while Unmarshalling path file.")
    fmt.Println(err)
    return false
  }

  // append the new data
  data.Paths = append(data.Paths, pathData.Paths...)

  // Marshal the updated yaml data
  updatedYaml, err := yaml.Marshal(&data)
  // Log if an error occurs.
  if err != nil {

    fmt.Println("Error occured while marhsalling path file.")
    fmt.Println(err)
    return false
  }

  // Writing the updated yaml data to the yaml file.
  err = os.WriteFile(filepath, updatedYaml, 0644)
  // Log if an error occurs
  if err != nil {

    fmt.Println("Error occured while writing to path file.")
    fmt.Println(err)
    return false
  }

  // safely exit the function.
  return true
}


// function to check whether a path is a directory and attach it to fsnotify
// WARN: Use only absolute paths. Would fail terribly otherwise.
func ifDirAttach (path string, watcher *fsnotify.Watcher) bool {

  info, error := os.Stat(path)
  if error != nil {

    fmt.Println("Error while pinging path to check for directory.")
    fmt.Println("Error Log:")
    fmt.Println(error)
    // BUG: Need to something more. Can't just return.
    // Send a notification and possible just revoke the application.
    return false
  }

  if !info.IsDir() {

    return false
  }

  watcher.Add(path)
  return true
}
