// FIXME: Add a type of error that can be returned by these functions.

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

    // TODO: Handle error
    return false
  }

  if (parsedUrl.Scheme == "" || parsedUrl.Host == "") {

    return false
  }

  if !strings.Contains(parsedUrl.Host, ".") {

    // BUG: well returning false would not specify why we return false.
    // Therefore, any indications to be made to the user would not be possible.
    return false
  }

  if !strings.HasPrefix(repoUrl, "https://github.com") {

    // BUG: Return an error code. Same as above.
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

    fmt.Println(err)
    // TODO: Handle the error.
    return false  
  }

  return true
}
