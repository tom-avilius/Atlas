/**
*
* @file: reponotifier.go
* @description: This file defines functions to check for changes in files and perform sync operations.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package reposync

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"tomavilius.in/atlas/internal/reporegistry"
)

/**
*
* @func dirList()
* @description: returns a list of child directories
*
**/

func dirList(path string) ([]string, bool) {

	// handling ~ paths
	if filePath, success := reporegistry.HandleHomeDirectory(path); success {

		path = filePath
	} else {

		return nil, false
	}

	// checking if the path exists
	if _, err := os.Stat(path); err != nil {

		// if an error occurs.
		fmt.Println("Error while pinging directory.")
		fmt.Println("Was about to check for child directories.")

		// if path also does not exist.
		if os.IsNotExist(err) {

			fmt.Println("The directory does not exist.")
		}
		fmt.Println("Error Log:")
		fmt.Println(err)

		return nil, false
	}

	// to store the list of directories
	var directories []string

	directories = append(directories, path)

	// traversing through the entire directory tree
	error := filepath.Walk(path, func(dir string, info os.FileInfo, err error) error {

		if err != nil {

			return err
		}

		// checking if the entity is a directory
		if info.IsDir() && path != dir {

			// appending it to the list of directories
			directories = append(directories, dir)
		}

		// successful function return
		return nil
	})

	// if an error occurs while traversing the directory tree
	// PERF: Use go's way of writing if statements.
	if error != nil {

		fmt.Println("Error occured while getting children directories.")
		fmt.Println("Error Log:")
		fmt.Println(error)
		return nil, false
	}

	// successful function ends successfulLy
	return directories, true
}

/**
*
* @func: attachFsnotify()
* @description: function to attach fsnotify to the dirlist
*
**/

func attachFsnotify(dirList []string) {

	// creating a file system watcher
	// PERF: Use go's way of writing if statements.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {

		fmt.Println("Error occured while creating watcher.")
		fmt.Println("Error Log:")
		fmt.Println(err)
		os.Exit(1)
	}
	defer watcher.Close()

	// go routineeeeeeees
	go func() {

		for {

			select {

			case event, ok := <-watcher.Events:
				if !ok {

					return
				}
				fmt.Println("event:", event)
				if event.Has(fsnotify.Write) {

					fmt.Println("Modified File")
					fmt.Println("File Path: " + event.Name)
					repoAdd(event.Name)
					repoCommit(event.Name, "Modify file: "+event.Name)
				}
				if event.Has(fsnotify.Remove) {

					// TODO: Remove from path file
					fmt.Println("File Removed")
					fmt.Println("File Path: " + event.Name)
					repoAdd(event.Name)
					repoCommit(event.Name, "Removed file: "+event.Name)
					WritePathData()
				}
				if event.Has(fsnotify.Create) {

					// TODO: If the new path is a directory then add it to watch list too.
					// TODO: Then update the path file.
					fmt.Println("New Path Created")
					fmt.Println("File Path: " + event.Name)
					repoAdd(event.Name)
					repoCommit(event.Name, "Created file: "+event.Name)
					ifDirAttach(event.Name, watcher)
					WritePathData()
				}
			case err, ok := <-watcher.Errors:
				if !ok {

					return
				}
				fmt.Println("Error: ")
				fmt.Println(err)
			}
		}
	}()

	// attach watcher to all directories and child directories.
	for _, dir := range dirList {

		err = watcher.Add(dir)
		if err != nil {

			fmt.Println(err)
			os.Exit(1)
		}
	}

	<-make(chan struct{})
}
