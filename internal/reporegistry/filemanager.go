
package reporegistry


import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)



// function the fills in the home directory when ~ is provided at the start
func handleHomeDirectory (filePath string) (string, bool) {

 if strings.HasPrefix(filePath, "~") {
    
    // getting the home directory
    homeDir, err := os.UserHomeDir()

    // Log if an error occurs
    if err != nil {

      fmt.Println("Error while getting the home directory.")
      fmt.Println("Error Log:")
      fmt.Println(err)
      return "", false
    }

    // otherwise concatenate home dir with the file path.
    filePath = strings.Join([]string{homeDir}, filePath[1:])
  }

  return filePath, true
}

// checkFileExist checks whether a file exists or not.
// returns false when it can not find the file or any error occurs.
func checkFileExist (filePath string) bool {

  // resolving for home dir
  if path, success := handleHomeDirectory(filePath); success {

    filePath = path
  } else {

    // unsuccessful 
    return false
  }
  
  // to check whether the file exists or not
  _, err := os.Stat(filePath)

  // return false when the file does not exist.
  if os.IsNotExist(err) {

    return false
  } else if err != nil { // Log if an error occured when checking for file existence

    fmt.Println("Error occured while pinging a file.")
    fmt.Println(err)
    return false
  }

  // safely exit function when no error occurs.
  return true;
}


// function to create a directory at the specified path.
// returns false if an error occurs
func createDir (dir string) bool {

  // create directory
  if direrr := os.MkdirAll(dir, os.ModePerm); direrr != nil {

    fmt.Println("Error occured while creating directory.")
    fmt.Println("Error Log:")
    fmt.Println(direrr)
    return false
  }

  // safely exit function
  return true
}

// createFile creates a new file at the specified location
// returns false if an error occurs.
// WARNING: If the file already exists then it will be formatted ( contents cleared ).
func createFile (filepath string) bool {

  // handling ~ paths
  if path, success := handleHomeDirectory(filepath); success {

    filepath = path
  } else {

    return false
  }

  // creating the file at the specified location
  file, err := os.Create(filepath)
  defer file.Close() 

  // Log if an error occurs.
  if err != nil {

    fmt.Println("Error occured while creating file.")
    fmt.Println(err)
    return false 
  }

  // safely exit function.
  return true
}


// writeYaml appends to a yaml file with repo information
// returns false if an error occurs.
func writeYaml (filepath string, repo Repository) bool {

  // to store yaml data.
  var data YamlData

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
  data.Repositories = append(data.Repositories, repo)

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


// deleteYaml deletes the specified yaml data from a yaml file.
// returns false if an error occurs.
func deleteYaml (filePath string, repoName string) bool {

  // to store the yaml data
  var data YamlData

  // reading the file
  content, err := os.ReadFile(filePath)
  // Log if an error occurs.
  if err != nil {

    fmt.Println("Error occured while reading a file.")
    fmt.Println(err)
    return false
  }

  // Unmarshal the read data.
  err = yaml.Unmarshal(content, &data)
  // Log if an error occurs
  if err != nil {

    fmt.Println("Error occured while Unmarshalling a file.")
    fmt.Println(err)
    return false
  }

  // amount of deleted data
  var flag int = 0;

  // to check for matching data
  for i := 0; i < len(data.Repositories); i++ {

    if data.Repositories[i].Name == repoName {

      // Remove the element at index i
      data.Repositories = append(data.Repositories[:i], data.Repositories[i+1:]...)
      // Since the slice shrinks, you need to decrement i to stay on the same index
      i--
      // increment flag
      flag++
    } 
  }

  // checking if no data data was deleted
  if flag == 0 {

    fmt.Println("Warning: No match found for " + repoName)
    return false;
  }

  // Marshal the updated data
  updatedYaml, err := yaml.Marshal(data)
  // Log if an error occurs.
  if err != nil {

    fmt.Println("Error occured while marshalling a file.")
    fmt.Println(err)
    return false
  }

  // Write the updated data to the yaml file.
  err = os.WriteFile(filePath, updatedYaml, 0644)
  // Log if an error occurs
  if err != nil {
    
    fmt.Println("Error occured while writing to a file.")
    fmt.Println(err)
    return false
  }

  // safely exit the function.
  fmt.Println("Success: Removed " +repoName +" from config.")
  return true
}

