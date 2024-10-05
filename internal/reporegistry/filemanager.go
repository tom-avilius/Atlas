package reporegistry

import (
	"fmt"
	"hash/adler32"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// checkFileExist checks whether a file exists or not.
func checkFileExist (filePath string) bool {

  if strings.HasPrefix(filePath, "~") {
    
    homeDir, err := os.UserHomeDir()

    if err != nil {

      fmt.Println("Error while getting the home directory.")
      fmt.Println(err)
      return false
    } else {

      filePath = strings.Join([]string{homeDir}, filePath[1:])
    }
  }

  _, err := os.Stat(filePath)

  if os.IsNotExist(err) {

    return false
  }

  if err != nil {

    fmt.Println("Error occured while pinging a file.")
    fmt.Println(err)
    return false
  }

  return true;
}

func createDir () bool {

  dir := "/home/tom-avilius/.config/atlas"

  _, err := os.Stat(dir)
  if os.IsNotExist(err) {

    err = os.MkdirAll(dir, os.ModePerm)
  }

  if err != nil {

    fmt.Println("Error occured while pinging directory.")
    fmt.Println(err)
    return false
  }

  return true
}

// createFile creates a new file at the specified location
func createFile (filepath string) bool {

  file, err := os.Create(filepath)
  defer file.Close() 

  if err != nil {

    fmt.Println("Error occured while creating file.")
    fmt.Println(err)
    return false 
  }

  return true
}

func writeYaml (filepath string, repo Repository) bool {

  var data yamlData

  content, err := os.ReadFile(filepath)
  if err != nil {

    fmt.Println("Error occured while reading a file.")
    fmt.Println(err)
    return false
  }

  err = yaml.Unmarshal(content, &data)
  if err != nil {

    fmt.Println("Error occured while Unmarshalling a file.")
    fmt.Println(err)
    return false
  }

  data.Repositories = append(data.Repositories, repo)

  updatedYaml, err := yaml.Marshal(&data)
  if err != nil {

    fmt.Println("Error occured while marhsalling a file.")
    fmt.Println(err)
    return false
  }

  err = os.WriteFile(filepath, updatedYaml, 0644)
  if err != nil {

    fmt.Println("Error occured while writing to a file.")
    fmt.Println(err)
    return false
  }

  return true
}

func deleteYaml (filePath string, repoName string) bool {

  var data yamlData

  content, err := os.ReadFile(filePath)
  if err != nil {

    fmt.Println("Error occured while reading a file.")
    fmt.Println(err)
    return false
  }

  err = yaml.Unmarshal(content, &data)
  if err != nil {

    fmt.Println("Error occured while Unmarshalling a file.")
    fmt.Println(err)
    return false
  }

  for i := 0 ; i < len(data.Repositories) ; i++ {

    if data.Repositories[i].Name == repoName {

      data.Repositories = append(data.Repositories[:i], data.Repositories[i+1:]...)
    } 
  }
  
  updatedYaml, err := yaml.Marshal(data)
  if err != nil {

    fmt.Println("Error occured while marshalling a file.")
    fmt.Println(err)
    return false
  }

  err = os.WriteFile(filePath, updatedYaml, 0644)
  if err != nil {
    
    fmt.Println("Error occured while writing to a file.")
    fmt.Println(err)
    return false
  }

  return true
}

