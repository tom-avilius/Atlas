
// Package repoRegistry provides data structures and interfaces for managing repositories.
// This module allows for adding, removing, and syncing repository information.
package reporegistry


import "time"


// INFO: Things could break because of this, just use absolute path if they do.
// Well they have not.
var ConfigDir string = "~/.config/atlas"
var ConfigFilePath string = "~/.config/atlas/config.yaml"

// Repository represents a git repository with its associated metadata.
// It includes the repository's name, URL, and timestamps for when it was added
// and last synced.
type Repository struct {
  
  Name     string    // Name of the repository
  Url      string    // URL of the repository
  Path     string
  AddedAt  time.Time // Timestamp when the repository was added
  LastSync time.Time // Timestamp of the last sync operation
}

// INFO: Not implemented so far

// Credentials holds user authentication details for accessing a repository.
// It includes the username and password required for access.
type Credentials struct {
  
  Username string // Username for accessing the repository
  Password string // Password for accessing the repository
}

// INFO: I Might not even have implemented this.

// RepositoryManager defines methods for managing repositories.
// It includes methods to add and remove repositories.
type RepositoryManager interface {

  // Add adds a repository with the specified credentials.
  Add(repo Repository, cred Credentials) error
  // Remove removes a repository by its name.
  Remove(repositoryName string) error
}

// RepoManager defines the location of the config file.
type RepoManager struct {

  ConfigFile string // the path to the YAML config file.
}

// yamlData defines the data to be entered into a yaml file
type YamlData struct {

  Repositories []Repository
}

