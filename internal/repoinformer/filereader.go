package repoinformer

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"tomavilius.in/atlas/internal/reporegistry"
)




func readYaml (filepath string) (*reporegistry.YamlData, bool) {

  // WARN: Expecting that the filepath is absolute.

  // if the file does not exist then exist the function
  if _, err := os.Stat(filepath); os.IsNotExist(err) {

    fmt.Println("Provided file does not exist.")
    fmt.Println("Error log:")
    fmt.Println(err)
    return nil, false
  }

  // read the yaml file
  content, err := os.ReadFile(filepath)
  // Log if an error occurs
  if err != nil {

    fmt.Println("Error occured while reading file.")
    fmt.Println("Error log:")
    fmt.Println(err)
    return nil, false
  }

  // empty data structure to hold yaml data
  var data reporegistry.YamlData

  // unmarshal the yaml data
  err = yaml.Unmarshal(content, &data)
  // log if an error occurs
  if err != nil {

    fmt.Println("Error occured while unmarshaling the yaml data.")
    fmt.Println("Error log:")
    fmt.Println(err)
    return nil, false
  }

  // safely exit the function
  return &data, true
}
