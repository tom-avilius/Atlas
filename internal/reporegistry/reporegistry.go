
package reporegistry


import "fmt"



// AddRepository adds a repository for atlas to manage.
// returns false if an error occurs.
func AddRepository (repo Repository, clonePath string) bool {

  // checking if the config file exists
  if !checkFileExist(ConfigFilePath) {

    // creating the file when it does not exist
    fmt.Println("\nConfig file does not exist.")
    createDir()
    fmt.Print("Creating config file.. ")
    createFile(ConfigFilePath)
    fmt.Print("Done.")
  }

  // validating the url.
  fmt.Print("\n\nValidating url.. ")
  isValidUrl := validateUrl(repo.Url);

  // cloning the repository
  if isValidUrl {

    fmt.Print("Done.")
    fmt.Println("\n\nProceeding to clone the repository..")
    fmt.Println()
    if cloneRepo(repo.Url, clonePath) {

      fmt.Print("\nUpdating the config.. ")
      if writeYaml(ConfigFilePath, repo) {

        fmt.Print("Done.")
        return true
      }
    }
  }

  // exit the function if it fails.
  fmt.Println("Error: Could not add repository.")
  return false
}


// DeleteRepository deletes the repository meaning atlas would no more manage it.
// returns false if an error occurs.
func DeleteRepository (onlyDeleteReference bool, repoName string) bool {

  // only delete the reference not the actual local folder.
  if onlyDeleteReference {

    return deleteYaml(ConfigFilePath, repoName);
  }

  return false
}

