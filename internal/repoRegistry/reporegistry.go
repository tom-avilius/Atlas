package reporegistry

import (
	"os"
	"strings"
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

func createFile (filepath string) bool {

  file, error := os.Create(filepath)
  defer file.Close() 

  if error != nil {

    return false 
  }

  return true
}
