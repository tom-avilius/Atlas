/**
*
* @file: filereader.go
* @description: This file provide functions to read yaml files..
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package repoinformer

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"tomavilius.in/atlas/internal/reporegistry"
)

/**
*
* @func: ReadYaml()
* @description: readYaml reades the yaml file and unmarshals it to YamlData structure
* return false if error occurs otherwise returns YamlData
*
**/
func ReadYaml(filepath string) (*reporegistry.YamlData, bool) {

	// handle ~ directory
	if path, success := reporegistry.HandleHomeDirectory(filepath); success {

		filepath = path
	} else {

		// unsuccessful
		return nil, false
	}

	// if the file does not exist then exist the function
	if _, err := os.Stat(filepath); os.IsNotExist(err) {

		fmt.Println("Provided file does not exist.")
		fmt.Println("Error log:")
		fmt.Println(err)
		return nil, false
	}

	// read the yaml file
	content, err := os.ReadFile(filepath)
	// PERF: Use go's way of writing if statements.
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
	// PERF: Use go's way of writing if statements.
	if err != nil {

		fmt.Println("Error occured while unmarshaling the yaml data.")
		fmt.Println("Error log:")
		fmt.Println(err)
		return nil, false
	}

	// safely exit the function
	return &data, true
}
