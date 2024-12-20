/**
*
* @file: reporegistry.go
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
	"path"
	"strings"
)

/**
*
* @func: function the fills in the home directory when ~ is provided at the start
*
**/
func HandleHomeDirectory(filePath string) (string, bool) {

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
		filePath = path.Join(homeDir, filePath[1:])
	}

	return filePath, true
}

// AddRepository adds a repository for atlas to manage.
// returns false if an error occurs.
func AddRepository(repo Repository, clonePath string) bool {

	// checking if the config file exists
	if !CheckFileExist(ConfigFilePath) {

		// creating the file when it does not exist
		fmt.Println("\nConfig file does not exist.")
		CreateDir(ConfigDir)
		fmt.Print("Creating config file.. ")
		CreateFile(ConfigFilePath)
		fmt.Print("Done.")
	}

	// validating the url.
	fmt.Print("\n\nValidating url.. ")
	isValidUrl := validateUrl(repo.Url)

	// checking if clone path exists
	if !CheckFileExist(clonePath) {

		// creating clone path if it does not exist
		fmt.Println("\nClone Path does not exist.")
		fmt.Println("Creating clone path..")
		CreateDir(clonePath)
	}

	// cloning the repository
	if isValidUrl {

		fmt.Print("Done.")
		fmt.Println("\n\nProceeding to clone the repository..")
		fmt.Println()
		if cloneRepo(repo.Url, path.Join(clonePath, repo.Name)) {

			fmt.Print("\nUpdating the config.. ")
			if WriteYaml(ConfigFilePath, repo) {

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
func DeleteRepository(onlyDeleteReference bool, repoName string) bool {

	// only delete the reference not the actual local folder.
	if onlyDeleteReference {

		return DeleteYaml(ConfigFilePath, repoName)
	}

	return false
}

// Replaced the entire config with a new empty one.
func ClearConfig() bool {

	if CreateFile(ConfigFilePath) {

		fmt.Println("\nCleared the config file.")
		return true
	} else {

		fmt.Println("\nCould not clear the config.")
		return false
	}
}
