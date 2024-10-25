package reposync

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)



func repoAdd(repoPath string) bool {

  // TODO: Use go's way of writing if statements.
  repo, err := git.PlainOpen(repoPath)
  if err != nil {

    fmt.Println("Error while opening repository.")
    fmt.Println("Error Log:")
    fmt.Println(err)
    return false
  }

  workTree, err := repo.Worktree()
  if err != nil {

    fmt.Println("Error while getting the Worktree.")
    fmt.Println("Error Log:")
    fmt.Println(err)
    return false
  }

  addOpts := &git.AddOptions {

    All: true,
  }

  err = workTree.AddWithOptions(addOpts)
  if err != nil {

    fmt.Println("Error while staging changes.")
    fmt.Println("Error Log:")
    fmt.Println(err)
  }

  fmt.Println("Staged changes.")
  return true
}
