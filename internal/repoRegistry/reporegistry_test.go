
package reporegistry

import "testing"

func TestCheckFileExist(t *testing.T) {
  
  // should return false
  result := checkFileExist("/path/to/file/that/does/notexist")
  if (result != false) {
    
    t.Errorf("Expected false, got %v", result)
  }

  // should return true
  result = checkFileExist("~/Projects/Atlas/tests/testConfig.YAML")
  if (result != true) {

    t.Errorf("Expected true, got %v", result)
  } 
}

