package reporegistry

import "testing"

func TestValidateUrl (t *testing.T) {

  // should return true
  result := validateUrl("https://github.com/tom-avilius")
  if result != true {

    t.Errorf("Expected true, got %v", result)
  }

  // should return false
  result = validateUrl("https:tom-avilius")
  if result != false {

    t.Errorf("Expected false, got %v", result)
  }
  result = validateUrl("//github.com/tom-avilius")
  if result != false {

    t.Errorf("Expected false, got %v", result)
  }
  result = validateUrl("https://tom-avilius")
  if result != false {

    t.Errorf("Expected false, got %v", result)
  }
  result = validateUrl("https:")
  if result != false {

    t.Errorf("Expected false, got %v", result)
  }
}
//
// func TestCloneRepo (t *testing.T) {
//
//   // should return true 
//   result := cloneRepo("https://github.com/tom-avilius/Atlas", "/home/tom-avilius/clones/atlas")
//   if result != true {
//
//     t.Errorf("Expected true, got %v", result)
//   }
// }
