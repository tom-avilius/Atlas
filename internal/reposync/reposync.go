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

  config, success := repoinformer.ReadYaml(reporegistry.ConfigFilePath);

  if !success {

    fmt.Println("Error while reading config file.")
    return false
  }

  for _, data := range config.Repositories {    

    dirData, success := dirList(data.Path)

    if !success {

      return false
    }

    var dirDataFormat reporegistry.PathData
    dirDataFormat.Paths = dirData

    if !writePathData(reporegistry.PathFilePath, dirDataFormat) {

      return false
    }
  }

  return true
}
