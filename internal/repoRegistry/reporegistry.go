
package reporegistry

func AddRepository (repo Repository, clonePath string) bool {

  // checking if the config file exists
  if !checkFileExist(configFilePath) {

    // creating the file when it does not exist
    createDir()
    createFile(configFilePath)
  }

  // validating the url.
  isValidUrl := validateUrl(repo.Url);

  if isValidUrl {

    if cloneRepo(repo.Url, clonePath) {

      return writeYaml(configFilePath, repo)
    }
  }

  return false
}

func DeleteRepository (onlyDeleteReference bool, repo Repository) bool {

  if onlyDeleteReference {

    return deleteYaml(configFilePath, repo.Name);
  }

  return true
}

