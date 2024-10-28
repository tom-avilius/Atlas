package atlasinstall

import (
	"fmt"
	"os/exec"
)



func InstallAtlas () bool {

  if !generateServiceFile() {

    return false
  }

  if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {

    fmt.Println("Error while running command: systemctl daemon-reload.")
    fmt.Println("Error Log:")
    fmt.Println(err)
    return false
  }

  if err := exec.Command("systemctl", "enable", "atlas").Run(); err != nil {

    fmt.Println("Error while running systemctl enable atlas")
    fmt.Println("Error Log:")
    fmt.Println(err)
    return false
  }

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
