/**
*
* @file: informer.go
* @description: This file exposes functions that read the config and print it.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package repoinformer

/**
*
* @func: DetailedRepositoryView()
* @description: DetailedRepositoryView shows formatted details from atlas config
* NOTE: It prints the entire config.
*
**/
func DetailedRepositoryView() bool {

	return listRepositories()
}

/**
*
* @func: ListRepositories()
* @description: ListRepositories() shows details from atlas config.
* NOTE: It prints only the repo names from atlas config.
*
**/
func ListRepositories() bool {

	return listOnlyRepositories()
}
