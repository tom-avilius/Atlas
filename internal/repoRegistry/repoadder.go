
// FIXME: Add a type of error that can be returned by these functions.

package reporegistry

import (
	"net/url"
	"strings"
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
