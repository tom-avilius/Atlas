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
  config, success := repoinformer.ReadYaml(reporegistry.ConfigFilePath);
  if !success {

    fmt.Println("Error while reading path file.")
    return false
  }
  fmt.Println("Done")

  // looping through the config data to extract path info 
  for _, data := range config.Repositories {    

    dirData, success := dirList(data.Path)

    if !success {

      return false
    }

    // to store path data
    var dirDataFormat reporegistry.PathData
    dirDataFormat.Paths = dirData

    // writing path data to the path file
    fmt.Println("\nWriting to path file..")
    if !writePathData(reporegistry.PathFilePath, dirDataFormat) {

      return false
    }
    fmt.Println("Done")

    attachFsnotify(dirData)
  }

  // successful execution
  return true
}
