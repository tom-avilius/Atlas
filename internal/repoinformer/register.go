package repoinformer

import (
	"fmt"
	"time"

	"tomavilius.in/atlas/internal/reporegistry"
)



func listRepositories () bool {

  data, success := readYaml(reporegistry.ConfigFilePath)

  // if read yaml was unsuccessful then simply return false
  if !success {

    return false
  }

  for i := 0; i < len(data.Repositories); i++ {

    fmt.Println()
    fmt.Println("Name: \t\t" +data.Repositories[i].Name)
    fmt.Println("Url: \t\t" +data.Repositories[i].Url)
    fmt.Println("Last Sync: \t" +data.Repositories[i].LastSync.Format(time.RFC850))
    fmt.Println("Added At: \t" +data.Repositories[i].AddedAt.Format(time.RFC850))
  }

  return success;
}

