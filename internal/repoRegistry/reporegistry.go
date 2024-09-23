package reporegistry

import (
	"fmt"
	"os"
	"strings"
	"gopkg.in/yaml.v3"
)

// checkFileExist checks whether a file exists or not.
func checkFileExist (filePath string) bool {

  if strings.HasPrefix(filePath, "~") {
    
    homeDir, err := os.UserHomeDir()

    if err != nil {

      // TODO: Handle this error
    } else {

      filePath = strings.Join([]string{homeDir}, filePath[1:])
    }
  }
  _, err := os.Stat(filePath)

  if os.IsNotExist(err) {

    return false
  }

  return true;
}

// createFile creates a new file at the specified location
func createFile (filepath string) bool {

  file, error := os.Create(filepath)
  defer file.Close() 

  // TODO: Handle all errors internally.
  if error != nil {

    return false 
  }

  return true
}

func writeYaml (filepath string, repo Repository) bool {

  var data yamlData

  content, err := os.ReadFile(filepath)
  if err != nil {

    // TODO: Handle errors internally.
    return false
  }

  err = yaml.Unmarshal(content, &data)
  if err != nil {

    // TODO: Handle error
    return false
  }

  fmt.Println(data)

  data.Repositories = append(data.Repositories, repo)

  updatedYaml, err := yaml.Marshal(&data)
  if err != nil {

    // TODO: Handle error
    return false
  }

  fmt.Println(data)
  fmt.Println(updatedYaml)

  err = os.WriteFile(filepath, updatedYaml, 0644)
  if err != nil {

    return false
  }

  return true
}
