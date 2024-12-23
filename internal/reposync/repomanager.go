/**
*
* @file: repomanager.go
* @description: This file defines functions to manipulate repositories and perform operations such as add, push and commit.
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

/**
*
* @func: repoAdd()
* @description: function to add changes to the working tree.
*
**/
func repoAdd(repoPath string) bool {

	// TODO: Use go's way of writing if statements.
	// openeing the repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {

		fmt.Println("Error while opening repository.")
		fmt.Println("Error Log:")
		fmt.Println(err)
		return false
	}

	// getting the current worktree
	workTree, err := repo.Worktree()
	if err != nil {

		fmt.Println("Error while getting the Worktree.")
		fmt.Println("Error Log:")
		fmt.Println(err)
		return false
	}

	// adding the repository changes to the work tree.
	addOpts := &git.AddOptions{

		All: true,
	}

	// PERF: Use go's way of writing if statements.
	err = workTree.AddWithOptions(addOpts)
	if err != nil {

		fmt.Println("Error while staging changes.")
		fmt.Println("Error Log:")
		fmt.Println(err)
	}

	fmt.Println("Staged changes.")
	return true
}

/**
*
* @func: repoCommit()
* @description: function to commit changes in a repository after adding them to the worktree.
**/
func repoCommit(repoPath, commitMessage string) bool {

	// PERF: Use go's way of writing if statements.
	// opening the git repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Println("Error while opening repository.")
		fmt.Println("Error Log:", err)
		return false
	}

	// PERF: Use go's way of writing if statements.
	// getting the current worktree
	workTree, err := repo.Worktree()
	if err != nil {
		fmt.Println("Error while getting the Worktree.")
		fmt.Println("Error Log:", err)
		return false
	}

	// PERF: Use go's way of writing if statements.
	// commiting the changes.
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

	// when the commit is successful, send out a success message.
	fmt.Println("Committed changes with hash:", commit.String())
	return true
}

/**
*
* @func: repoPush()
* @description: function to push changes to the remote repository.
*
**/
func repoPush(repoPath, username, password string) bool {

	// PERF: Use go's way of writing if statements.
	// opening the git repository.
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Println("Error while opening repository.")
		fmt.Println("Error Log:", err)
		return false
	}

	// PERF: Use go's way of writing if statements.
	// pushing the changes to the remote repository.
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

	// print success message when function successful.
	fmt.Println("Pushed changes to remote repository.")
	return true
}
