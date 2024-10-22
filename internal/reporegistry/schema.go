        
// Package repoRegistry provides data structures and interfaces for managing repositories.
// This module allows for adding, removing, and syncing repository information.
package reporegistry


import "time"


// FIXME: Provide a dynamic path for home-dir
// config File path
var ConfigDir = "/home/tom_avilius/.config/atlas"
var ConfigFilePath string = "/home/tom_avilius/.config/atlas/config.yaml"

// Repository represents a git repository with its associated metadata.
// It includes the repository's name, URL, and timestamps for when it was added
// and last synced.
type Repository struct {
  
  Name     string    // Name of the repository
  Url      string    // URL of the repository
  AddedAt  time.Time // Timestamp when the repository was added
  LastSync time.Time // Timestamp of the last sync operation
}

// Credentials holds user authentication details for accessing a repository.
// It includes the username and password required for access.
type Credentials struct {
  
  Username string // Username for accessing the repository
  Password string // Password for accessing the repository
}

// RepositoryManager defines methods for managing repositories.
// It includes methods to add and remove repositories.
type RepositoryManager interface {

  // TODO: Instead of using the standard error interface, use an advanced and modified version from a local module instead.

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

