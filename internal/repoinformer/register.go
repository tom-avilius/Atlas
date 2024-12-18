/**
*
* @file: register.go
* @description: This file defines functions that read the config and print it.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package repoinformer

import (
	"fmt"
	"time"

	"tomavilius.in/atlas/internal/reporegistry"
)

/**
*
* @func: listRepositories()
* @description: listRepositories lists the complete atlas config after formatting it.
*
**/
func listRepositories() bool {

	// read the config file
	data, success := ReadYaml(reporegistry.ConfigFilePath)
	// if read yaml was unsuccessful then simply return false
	// PERF: Use go's way of writing if statements.
	if !success {

		fmt.Println("Could not read atlas config file. Maybe you have not provided any repository for atlas to watch yet.")
		return false
	}

	// format the config to display
	for i := 0; i < len(data.Repositories); i++ {

		fmt.Println()
		fmt.Println("Name: \t\t" + data.Repositories[i].Name)
		fmt.Println("Url: \t\t" + data.Repositories[i].Url)
		fmt.Println("Path: \t\t" + data.Repositories[i].Path)
		fmt.Println("Last Sync: \t" + data.Repositories[i].LastSync.Format(time.RFC850))
		fmt.Println("Added At: \t" + data.Repositories[i].AddedAt.Format(time.RFC850))
	}

	return success
}

/**
*
* @func: listOnlyRepositories()
* @description: to only list Repositories from atlas config.
*
**/
func listOnlyRepositories() bool {

	// read the config file
	data, success := ReadYaml(reporegistry.ConfigFilePath)
	// if read yaml was unsuccessful then simply return false
	// PERF: Use go's way of writing if statuements.
	if !success {

		fmt.Println("Could not read atlas config file. Maybe you have not added any repository yet for atlas to watch.")
		return false
	}

	// display Repositories
	fmt.Println() // empty line..
	for i := 0; i < len(data.Repositories); i++ {

		// print the repo name
		fmt.Println(data.Repositories[i].Name)
	}

	return success
}
