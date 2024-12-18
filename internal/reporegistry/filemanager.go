/**
*
* @file: filemanager.go
* @description: This file defines functions to manipulate YAML files.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package reporegistry

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// TODO: Port the ReadYaml function I found elsewhere to this file.

/**
*
* @function CheckFileExist()
* @description: checkFileExist checks whether a file exists or not.
* returns false when it can not find the file or any error occurs.
*
**/
func CheckFileExist(filePath string) bool {

	// resolving for home dir
	if path, success := HandleHomeDirectory(filePath); success {

		filePath = path
	} else {

		// unsuccessful
		return false
	}

	// to check whether the file exists or not
	// PERF: Use go's way of writing if statements.
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {

		return false
	} else if err != nil { // Log if an error occured when checking for file existence

		fmt.Println("Error occured while pinging a file.")
		fmt.Println(err)
		return false
	}

	// safely exit function when no error occurs.
	return true
}

/**
*
* @function CreateDir()
* @description: function to create a directory at the specified path.
* returns false if an error occurs
*
**/
func CreateDir(dir string) bool {

	// resolving ~ directory
	if path, success := HandleHomeDirectory(dir); success {

		dir = path
	} else {

		return false
	}

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

/**
*
* @function CreateFile()
* @description: createFile creates a new file at the specified location
* returns false if an error occurs.
*
**/
// WARNING: If the file already exists then it will be formatted ( contents cleared ).
func CreateFile(filepath string) bool {

	// handling ~ paths
	if path, success := HandleHomeDirectory(filepath); success {

		filepath = path
	} else {

		// unsuccessful
		return false
	}

	// creating the file at the specified location
	file, err := os.Create(filepath)
	defer file.Close()
	// PERF: Use go's way of writing if statements.
	if err != nil {

		fmt.Println("Error occured while creating file.")
		fmt.Println(err)
		return false
	}

	// safely exit function.
	return true
}

/**
*
* @function WriteYaml()
* @description: writeYaml appends to a yaml file with repo information
* returns false if an error occurs.
*
**/
func WriteYaml(filepath string, repo Repository) bool {

	// handling ~ paths
	if path, success := HandleHomeDirectory(filepath); success {

		filepath = path
	} else {

		// unsuccessful
		return false
	}

	// to store yaml data.
	var data YamlData

	// read the yaml file
	// PERF: Use go's way of writing if statements.
	content, err := os.ReadFile(filepath)
	if err != nil {

		fmt.Println("Error occured while reading a file.")
		fmt.Println(err)
		return false
	}

	// Unmarshal the yaml data
	// PERF: Use go's way of writing if statements.
	err = yaml.Unmarshal(content, &data)
	if err != nil {

		fmt.Println("Error occured while Unmarshalling a file.")
		fmt.Println(err)
		return false
	}

	// append the new data
	data.Repositories = append(data.Repositories, repo)

	// Marshal the updated yaml data
	// PERF: Use go's way of writing if statements.
	updatedYaml, err := yaml.Marshal(&data)
	if err != nil {

		fmt.Println("Error occured while marhsalling a file.")
		fmt.Println(err)
		return false
	}

	// Writing the updated yaml data to the yaml file.
	// PERF: Use go's way of writing if statements.
	err = os.WriteFile(filepath, updatedYaml, 0644)
	if err != nil {

		fmt.Println("Error occured while writing to a file.")
		fmt.Println(err)
		return false
	}

	// safely exit the function.
	return true
}

/**
*
* @function DeleteYaml()
* @description: deleteYaml deletes the specified yaml data from a yaml file.
*	returns false if an error occurs.
*
**/
func DeleteYaml(filePath string, repoName string) bool {

	// handling ~ paths
	if path, success := HandleHomeDirectory(filePath); success {

		filePath = path
	} else {

		// unsuccessful
		return false
	}

	// to store the yaml data
	var data YamlData

	// reading the file
	// PERF: Use go's way of writing if statements.
	content, err := os.ReadFile(filePath)
	if err != nil {

		fmt.Println("Error occured while reading a file.")
		fmt.Println(err)
		return false
	}

	// Unmarshal the read data.
	// PERF: Use go's way of writing if statements.
	err = yaml.Unmarshal(content, &data)
	if err != nil {

		fmt.Println("Error occured while Unmarshalling a file.")
		fmt.Println(err)
		return false
	}

	// amount of deleted data
	var flag int = 0

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
		return false
	}

	// Marshal the updated data
	// PERF: Use go's way of writing if statements.
	updatedYaml, err := yaml.Marshal(data)
	if err != nil {

		fmt.Println("Error occured while marshalling a file.")
		fmt.Println(err)
		return false
	}

	// Write the updated data to the yaml file.
	// PERF: Use go's way of writing if statements.
	err = os.WriteFile(filePath, updatedYaml, 0644)
	if err != nil {

		fmt.Println("Error occured while writing to a file.")
		fmt.Println(err)
		return false
	}

	// safely exit the function.
	fmt.Println("Success: Removed " + repoName + " from config.")
	return true
}
