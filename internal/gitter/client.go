package gitter

import (
	"context"
	"pierflow/internal/eventer"
	"sync"
)

type GitClient interface {
	// Clone a git repository to the specified path.
	Clone(ctx context.Context, o *CloneOptions, messager eventer.Messager)

	// BranchList lists branches in the specified repository path.
	BranchList(ctx context.Context, options *BranchOptions, messager eventer.Messager)

	// Checkout a specific branch in the given repository path.
	Checkout(o *CheckoutOptions, messager eventer.Messager)

	// Pull the latest changes from the remote repository.
	Pull(ctx context.Context, o *PullOptions, messager eventer.Messager)
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
