/**
*
* @file: reposync.go
* @description: This file exposes functions to manage the syncing of local repositories.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package reposync

import (
	"fmt"
	"strings"

	"tomavilius.in/atlas/internal/repoinformer"
	"tomavilius.in/atlas/internal/reporegistry"
)

/**
*
* @func: WritePathData()
* @description: function to write path data to path file
*
**/
func WritePathData() bool {

	// creating path file if it does not exist
	if success := createPathFile(); !success {

		return false
	}

	// reading the config file
	// FIXME: You are reading the config file here not the path file.
	fmt.Println("\nReading Path file..")
	// PERF: Use go's way of writing if statement
	config, success := repoinformer.ReadYaml(reporegistry.ConfigFilePath)
	if !success {

		fmt.Println("Error while reading path file.")
		return false
	}
	fmt.Println("Done")

	// to store path data
	var dirDataFormat reporegistry.PathData

	// looping through the config data to extract path info
	for _, data := range config.Repositories {

		// create list of child directories.
		// PERF: Use go's way of writing if statements.
		dirData, success := dirList(data.Path)
		if !success {

			return false
		}

		dirDataFormat.Paths = append(dirDataFormat.Paths, dirData...)

		// writing path data to the path file
		fmt.Println("\nWriting to path file..")
		if !writePathData(reporegistry.PathFilePath, dirDataFormat) {

			return false
		}
		fmt.Println("Done")
	}

	// successful execution
	return true
}

// TODO: Create a function to take in official paths without the child paths
// and use it to map a child path back to its father...

/**
*
* @func: mapBackChildPath()
* @description: function to map a child path back to its parent path
*
**/
func mapBackChildPath(path string) (string, bool) {

	// reading the config file.
	// WARN: Ensure that the config file exists.
	if configData, success := repoinformer.ReadYaml(reporegistry.ConfigFilePath); success {

		// looping over the config repository data to get repo paths.
		for _, data := range configData.Repositories {

			// checking if the repo path begins with the provided child path
			if strings.Contains(path, data.Path) {

				return data.Path, true
			}
		}
	} else {

		// could not access the config file.
		fmt.Println("Could not read the atlas config file; was trying to map child paths back to their parents.")
	}

	// return in case of failure.
	return "", false
}

/**
*
* @func: AttachFsNotify()
* @description: to add all dirpaths to fsnotify
*
**/
func AttachFsNotify() bool {

	// prompt action
	fmt.Println("\nAttaching fsnotify..")

	// reading the config file
	fmt.Println("Reading Path file..")
	// PERF: Use go's way of writing if statements.
	config, success := repoinformer.ReadYaml(reporegistry.ConfigFilePath)
	if !success {

		fmt.Println("Error while reading path file.")
		return false
	}
	fmt.Println("Done")

	// to store path data
	var dirDataFormat reporegistry.PathData

	// looping through the config data to extract path info
	for _, data := range config.Repositories {

		// finding child directories of each local repository
		// then add them to the list of all child directories containing
		// child directories of all local repositories.
		// PERF: Use go's way of writing if statements.
		dirData, success := dirList(data.Path)
		if !success {

			return false
		}

		// the actual appending.
		dirDataFormat.Paths = append(dirDataFormat.Paths, dirData...)
	}

	// attaching fsnotify
	attachFsnotify(dirDataFormat.Paths)

	// successful
	return true
}
