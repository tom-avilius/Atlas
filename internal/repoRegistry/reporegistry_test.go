
package reporegistry

import "testing"

func TestCheckFileExist (t *testing.T) {
  
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

func TestCreateFile (t *testing.T) {

  //should return true
  result := createFile("/home/tom-avilius/Projects/Atlas/tests/hello.yaml")
  if result != true {

    t.Errorf("Expected true, got %v", result)
  }

  // should return false 
  result = createFile("/////asdasd")
  if result != false {

    t.Errorf("Expected false, got %v", result)
  }
}
