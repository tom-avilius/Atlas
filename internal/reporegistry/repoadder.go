
package reporegistry

import (
	
  "fmt"
	"net/url"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)


func validateUrl (repoUrl string) bool {

  parsedUrl, err := url.ParseRequestURI(repoUrl);
  if err != nil {

    fmt.Println("Error occured while parsing url.")
    fmt.Println(err)
    return false
  }

  if (parsedUrl.Scheme == "" || parsedUrl.Host == "") {

    return false
  }

  if !strings.Contains(parsedUrl.Host, ".") {

    fmt.Println("Invalid Url Provided.")
    return false
  }

  if !strings.HasPrefix(repoUrl, "https://github.com") {

    fmt.Println("Error: Not a github link.")
    return false 
  }
  
  return true
}


func cloneRepo (repoUrl string, clonePath string) bool {

  _, err := git.PlainClone(clonePath, false, &git.CloneOptions{

    URL: repoUrl,
    Progress: os.Stdout,
  })

  if err != nil {

    fmt.Println("Error occured while cloning the repository.")
    fmt.Println(err)
    return false  
  }

  return true
}

