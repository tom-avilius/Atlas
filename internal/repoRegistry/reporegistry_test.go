package reporegistry

import (
	"testing"
	"time"
)


func TestAddRepository (t *testing.T) {

  repository := Repository {

    Name: "splash",
    Url: "https://github.com/tom-avilius/splash",
    AddedAt: time.Now(),
    LastSync: time.Time{},
  }

  // should return true
  result := AddRepository(repository, "/home/tom-avilius/Projects/" + repository.Name )
  if (result != true) {

    t.Errorf("Expected true, got %v", false)
  }
}
