package gitter

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"pierflow/internal/logger"
	"sync"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type CheckoutOptions struct {
	Branch string `json:"branch"`
	Place  string `json:"place"`
}

type BranchOptions struct {
	Refresh bool   `json:"refresh"`
	User    string `json:"user"`
	Token   string `json:"token"`
	Prune   bool   `json:"prune"`
}

type PullOptions struct {
	User   string `json:"user"`
	Token  string `json:"token"`
	GitUrl string `json:"gitUrl"`
}

type GitClient interface {
	Clone(ctx context.Context, user, token, repoUrl, path string) (string, error)
	BranchList(ctx context.Context, path string, options *BranchOptions) (string, []Branch, error)
	CheckoutBranch(path string, o *CheckoutOptions) (*Branch, error)
	Pull(ctx context.Context, path string, o *PullOptions) (string, *Branch, error)
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

func (g *gitClient) Clone(ctx context.Context, user, token, repoUrl, path string) (string, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	repositoryPath := filepath.Join(g.basePath, path)

	// Error if repositoryPath is existing
	if _, err := os.Stat(repositoryPath); err == nil {
		return "", errors.New("repository already exists")
	}
	logger.Infof("cloning repository from %s to %s", repoUrl, repositoryPath)

	progressor := newProgressor()

	_, err := git.PlainCloneContext(ctx, repositoryPath, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: progressor,
		Auth: &http.BasicAuth{
			Username: user,
			Password: token,
		},
	})

	if err != nil {
		return "", err
	}

	// Implement the logic to clone a git repository
	// This is a placeholder implementation
	return progressor.String(), nil
}
