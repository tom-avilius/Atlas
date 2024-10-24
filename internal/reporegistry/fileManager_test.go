 package reporegistry

import (
	"testing"
	"time"
)

func TestCheckFileExist (t *testing.T) {
  
  // should return false
  result := CheckFileExist("/path/to/file/that/does/notexist")
  if (result != false) {
    
    t.Errorf("Expected false, got %v", result)
  }

  // should return true
  result = CheckFileExist("~/Projects/Atlas/tests/testConfig.YAML")
  if (result != true) {

    t.Errorf("Expected true, got %v", result)
  } 
}

func TestCreateFile (t *testing.T) {

  //should return true
  result := CreateFile("/home/tom-avilius/Projects/Atlas/tests/hello.yaml")
  if result != true {

    t.Errorf("Expected true, got %v", result)
  }

  // should return false 
  result = CreateFile("/////asdasd")
  if result != false {

    t.Errorf("Expected false, got %v", result)
  }
}

func TestWriteYaml (t *testing.T) {

  var repository = Repository {
    
    Name: "lo",
    Url: "eees",
    AddedAt: time.Now(),
    LastSync: time.Now(),
  }

  // should return true
  result := WriteYaml("/home/tom-avilius/Projects/Atlas/tests/testConfig.yaml", repository)
  if result != true {

    t.Errorf("Expected true, got %v", result)
  }

  // should return false
  result = WriteYaml("////sdads", repository)
  if result != false {

    t.Errorf("Expected false, got %v", result)
  }
}

func TestDeleteYaml (t *testing.T) {

  // should return true
  result := DeleteYaml("/home/tom-avilius/Projects/Atlas/tests/testConfig.yaml", "lo")
  if result != true {

    t.Errorf("Expected true, got %v", result)
  }

  // should return false
  result = DeleteYaml("////sdads", "lo")
  if result != false {

    t.Errorf("Expected false, got %v", result)
  }
}

