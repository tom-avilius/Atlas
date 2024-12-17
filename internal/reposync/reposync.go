package reposync

import (
	"fmt"

	"tomavilius.in/atlas/internal/repoinformer"
	"tomavilius.in/atlas/internal/reporegistry"
)

// function to write path data to path file
func WritePathData() bool {

	// creating path file if it does not exist
	if success := createPathFile(); !success {

		return false
	}

	// reading the config file
	fmt.Println("\nReading Path file..")
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

// function to map a child path back to its parent path
func mapBackChildPath(path string) string { return path }

// to add all dirpaths to fsnotify
func AttachFsNotify() bool {

	fmt.Println("\nAttaching fsnotify..")

	// reading the config file
	fmt.Println("Reading Path file..")
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

		dirData, success := dirList(data.Path)
		if !success {

			return false
		}

		dirDataFormat.Paths = append(dirDataFormat.Paths, dirData...)
	}

	// attaching fsnotify
	attachFsnotify(dirDataFormat.Paths)

	// successful
	return true
}
