/**
*
* @file: install.go
* @description: This file brings together code to create systemd service file, enable service and reload systemctl and to start the service.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package atlasinstall

import (
	"fmt"
	"os/exec"
)

/**
*
* @func: InstallAtlas()
* @description: Brings together code to create, reload, start and enable the atlas service.
*
**/
func InstallAtlas() bool {

	// creating the systemd service file.
	if !generateServiceFile() {

		return false
	}

	// reloading the systemd daemon.
	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {

		fmt.Println("Error while running command: systemctl daemon-reload.")
		fmt.Println("Error Log:")
		fmt.Println(err)
		return false
	}

	// enabling the atlas service.
	if err := exec.Command("systemctl", "enable", "atlas").Run(); err != nil {

		fmt.Println("Error while running systemctl enable atlas")
		fmt.Println("Error Log:")
		fmt.Println(err)
		return false
	}

	// starting the atlas service.
	if err := exec.Command("systemctl", "start", "atlas").Run(); err != nil {

		fmt.Println("Error while running systemctl start atlas.")
		fmt.Println("Error Log:")
		fmt.Println(err)
		return false
	}

	// TODO: Maybe include a notification
	fmt.Println("Systemd service created and started successfully.")
	return true
}
