package reposync

import (
	"fmt"

	"tomavilius.in/atlas/internal/repoinformer"
	"tomavilius.in/atlas/internal/reporegistry"
)

// function to write path data
func WritePathData() bool {

  if success := createPathFile(); !success {

    return false
  }

  fmt.Println("\nReading Path file..")
  config, success := repoinformer.ReadYaml(reporegistry.ConfigFilePath);
  if !success {

    fmt.Println("Error while reading path file.")
    return false
  }
  fmt.Println("Done")

  for _, data := range config.Repositories {    

    dirData, success := dirList(data.Path)

    if !success {

      return false
    }

    var dirDataFormat reporegistry.PathData
    dirDataFormat.Paths = dirData

    fmt.Println("\nWriting to path file..")
    if !writePathData(reporegistry.PathFilePath, dirDataFormat) {

      return false
    }
    fmt.Println("Done")
  }

  fmt.Println("\nSyncing Completed.")
  return true
}
