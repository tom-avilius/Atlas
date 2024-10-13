
package reporegistry


import (
	
  "fmt"
	"net/url"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)



// validateUrl checks if the url is valid.
// returns false if the url is not valid or an error occurs.
// url should be of a github repository.
func validateUrl (repoUrl string) bool {

  // parsing the url
  parsedUrl, err := url.ParseRequestURI(repoUrl);
  // Log if an error occurs
  if err != nil {

    fmt.Println("Error occured while parsing url.")
    fmt.Println(err)
    return false
  }

  // check for its validity
  if (parsedUrl.Scheme == "" || parsedUrl.Host == "") {

    return false
  }

  // check for Host
  if !strings.Contains(parsedUrl.Host, ".") {

    fmt.Println("Invalid Url Provided.")
    return false
  }

  // check for github url
  if !strings.HasPrefix(repoUrl, "https://github.com") {

    fmt.Println("Error: Not a github link.")
    return false 
  }
  
  // safely exist the function.
  return true
}


// cloneRepo clones the given repository
// return false if an error occurs
func cloneRepo (repoUrl string, clonePath string) bool {

  // performing plain clone of the repository.
  _, err := git.PlainClone(clonePath, false, &git.CloneOptions{

    URL: repoUrl,
    Progress: os.Stdout,
  })

  // Log if an error occurs. 
  if err != nil {

    fmt.Println("Error occured while cloning the repository.")
    fmt.Println(err)
    return false  
  }

  // safely exit the function
  return true
}

