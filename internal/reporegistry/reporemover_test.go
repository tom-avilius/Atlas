
package reporegistry


import "testing"

func TestRepositoryRemover (t *testing.T) {

  //should return true 
  result := removeRepository("lo", true);
  if result != true {

    t.Errorf("Expected true, got %v", result)
  }
}

