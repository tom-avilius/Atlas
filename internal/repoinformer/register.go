package repoinformer

import (
	"fmt"

	"tomavilius.in/atlas/internal/reporegistry"
)



func listRepositories () bool {

  data, success := readYaml(reporegistry.ConfigFilePath)

  if success {

    fmt.Println(data);
  }

  return success;
}

