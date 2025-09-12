package gitter

import (
	"sync"

	"github.com/blueskyfish/pierflow/internal/eventer"
)

type GitClient interface {
	// Clone a git repository to the specified path.
	//
	// On success, it sends the current branch information through the messager.
	Clone(o *CloneOptions, messager eventer.Messager)

	// BranchList lists branches in the specified repository path.
	//
	// On success, it sends the list of branches through the messager.
	BranchList(options *BranchOptions, messager eventer.Messager)

	// Checkout a specific branch in the given repository path.
	//
	// On success, it sends the current branch information through the messager.
	Checkout(o *CheckoutOptions, messager eventer.Messager)

	// Pull the latest changes from the remote repository.
	//
	// On success, it sends the current branch information through the messager.
	Pull(o *PullOptions, messager eventer.Messager)
}

type gitClient struct {
	basePath string
	mutex    sync.RWMutex
}

// NewGitClient initializes a new GitClient with the specified base path.
//
// The parameter `basePath` is the directory where git repositories will be cloned and is required for the client to function properly.
func NewGitClient(basePath string) GitClient {
	return &gitClient{
		basePath: basePath,
		mutex:    sync.RWMutex{},
	}
}
