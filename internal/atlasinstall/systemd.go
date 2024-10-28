package atlasinstall

import (
	"fmt"
	"os"
)

// TODO: Move all systemd file functions to a single file..
func generateServiceFile () bool {

  execPath, err := os.Executable();
  if err != nil {

    fmt.Println("Error occured while getting executable path.")
    fmt.Println("Error Log:")
    fmt.Println(err)
    return false
  }

  fileContent := fmt.Sprintf(`[Unit]
  Description=Atlas Daemon
  After=network.target

  [Service]
  Type=simple
  ExecStart=%s start

  [Install]
  WantedBy=multi-user.target
  `, execPath)

  servicePath := "/etc/systemd/system/atlas.service"
  file, err := os.Create(servicePath)
  if err != nil {
    
    fmt.Println("Error occured while creating atlas service path.")
    fmt.Println("Error Log:")
    fmt.Println(err)
    return false
  }
  defer file.Close()

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
