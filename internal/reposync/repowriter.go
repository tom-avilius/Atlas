package reposync

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"tomavilius.in/atlas/internal/reporegistry"
)

// function to create a path file which stores the various paths that must be added to fsnotify
func createPathFile() bool {

  // if the path file does not exist, create it
  if !reporegistry.CheckFileExist(reporegistry.PathFilePath) {
    
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
  }

  // successful
  return true
} 


// function to write path data to the path file.
// returns false if an error occurs.
func writePathData (filepath string, pathData []string) bool {

  // handling ~ paths
  if path, success := reporegistry.HandleHomeDirectory(filepath); success {

    filepath = path
  } else {
    
    // unsuccessful 
    return false
  }

  // to store yaml data.
  var data reporegistry.YamlData

  // read the yaml file
  content, err := os.ReadFile(filepath)
  // Log if an error occurs
  if err != nil {

    fmt.Println("Error occured while reading a file.")
    fmt.Println(err)
    return false
  }

  // Unmarshal the yaml data
  err = yaml.Unmarshal(content, &data)
  // Log if an error occurs
  if err != nil {

    fmt.Println("Error occured while Unmarshalling a file.")
    fmt.Println(err)
    return false
  }

  // append the new data
   = append(data.Reposito, pathData)

  // Marshal the updated yaml data
  updatedYaml, err := yaml.Marshal(&data)
  // Log if an error occurs.
  if err != nil {

    fmt.Println("Error occured while marhsalling a file.")
    fmt.Println(err)
    return false
  }

  // Writing the updated yaml data to the yaml file.
  err = os.WriteFile(filepath, updatedYaml, 0644)
  // Log if an error occurs
  if err != nil {

    fmt.Println("Error occured while writing to a file.")
    fmt.Println(err)
    return false
  }

  // safely exit the function.
  return true
}

