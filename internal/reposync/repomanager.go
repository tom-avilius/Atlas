/**
*
* @file: repomanager.go
* @description: This file defines functions to manipulate YAML files.
*
* @author: tom avilius <tomavilius@tomavilius.in>
* @license: MIT license
* @package: Atlas v0.0.1 development
*
**/

package reposync

// WARN: Test beviour of commit and push.
// TODO: Somewhere store a personal access token and username and email for commiting and pushing to work.

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
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

	addOpts := &git.AddOptions{

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

// Function to commit changes
func repoCommit(repoPath, commitMessage string) bool {

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Println("Error while opening repository.")
		fmt.Println("Error Log:", err)
		return false
	}

	workTree, err := repo.Worktree()
	if err != nil {
		fmt.Println("Error while getting the Worktree.")
		fmt.Println("Error Log:", err)
		return false
	}

	commit, err := workTree.Commit(commitMessage, &git.CommitOptions{
		Author: &object.Signature{
			//TODO: Add user name and email
			Name:  "Your Name",      // Replace with the author's name
			Email: "your.email@xyz", // Replace with the author's email
			When:  time.Now(),
		},
	})
	if err != nil {
		fmt.Println("Error while committing changes.")
		fmt.Println("Error Log:", err)
		return false
	}

	fmt.Println("Committed changes with hash:", commit.String())
	return true
}

// Function to push changes
func repoPush(repoPath, username, password string) bool {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Println("Error while opening repository.")
		fmt.Println("Error Log:", err)
		return false
	}

	err = repo.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: username, // Username for GitHub or other Git service
			Password: password, // Password or personal access token
		},
	})
	if err != nil {
		fmt.Println("Error while pushing changes.")
		fmt.Println("Error Log:", err)
		return false
	}

	fmt.Println("Pushed changes to remote repository.")
	return true
}
