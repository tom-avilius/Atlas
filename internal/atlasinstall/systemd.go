/**
*
* @file: systemd.go
* @description: This file creates a systemd service file for atlas.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package atlasinstall

import (
	"fmt"
	"os"
)

/**
*
* @func: generateServiceFile()
* @description: function to generate atlas service file to run the start command on login.
*
**/
// TODO: Move all systemd file functions to a single file..
func generateServiceFile() bool {

	// getting the executable path directory for atals binary.
	// PERF: Use go's way of writing if statements.
	execPath, err := os.Executable()
	if err != nil {

		fmt.Println("Error occured while getting executable path.")
		fmt.Println("Error Log:")
		fmt.Println(err)
		return false
	}

	// service file content for atlas.
	fileContent := fmt.Sprintf(`[Unit]
  Description=Atlas Daemon
  After=network.target

  [Service]
  Type=simple
  ExecStart=%s start

  [Install]
  WantedBy=multi-user.target
  `, execPath)

	// path where the service file is to be place.
	// TEST: Should this be somewhere else in a different file... or have global scope?
	servicePath := "/etc/systemd/system/atlas.service"

	// creating the systemd service file at the service path.
	// PERF: Use go's way of writing if statements.
	file, err := os.Create(servicePath)
	if err != nil {

		fmt.Println("Error occured while creating atlas service path.")
		fmt.Println("Error Log:")
		fmt.Println(err)
		return false
	}
	defer file.Close()

	// writing to the service file the service content.
	// PERF: Use go's way of writing if statements.
	_, err = file.WriteString(fileContent)
	if err != nil {

		fmt.Println("Error occured while writing to atlas service file.")
		fmt.Println("Error Log:")
		fmt.Println(err)
		return false
	}

	// TODO: Maybe show a notification
	fmt.Println("Atlas service file created.")
	return true
}
