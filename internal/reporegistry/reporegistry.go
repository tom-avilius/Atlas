
package reporegistry

import "fmt"

func AddRepository (repo Repository, clonePath string) bool {

  // checking if the config file exists
  if !checkFileExist(configFilePath) {

    // creating the file when it does not exist
    fmt.Println("\nConfig file does not exist.")
    createDir()
    fmt.Print("Creating config file.. ")
    createFile(configFilePath)
    fmt.Print("Done.")
  }

  fmt.Print("\n\nValidating url.. ")
  // validating the url.
  isValidUrl := validateUrl(repo.Url);

  if isValidUrl {

    fmt.Print("Done.")
    fmt.Println("\n\nProceeding to clone the repository..")
    fmt.Println()
    if cloneRepo(repo.Url, clonePath) {

      fmt.Print("\nUpdating the config.. ")
      if writeYaml(configFilePath, repo) {

        fmt.Print("Done.")
        return true
      }
    }
  }

  fmt.Println("Error: Could not add repository.")
  return false
}

func DeleteRepository (onlyDeleteReference bool, repo Repository) bool {

  if onlyDeleteReference {

    return deleteYaml(configFilePath, repo.Name);
  }

  fmt.Println("Error: Could not delete repository.")
  return false
}

