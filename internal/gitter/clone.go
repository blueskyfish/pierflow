package gitter

import (
	"context"
	"os"
	"path/filepath"
	"pierflow/internal/eventer"
	"pierflow/internal/logger"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type CloneOptions struct {
	User    string
	Token   string
	RepoUrl string
	Path    string
}

func (g *gitClient) Clone(ctx context.Context, o *CloneOptions, messager eventer.Messager) {
	go g.runClone(ctx, o, messager)
}

// runClone performs the actual cloning operation in a separate goroutine.
//
// It locks the mutex to ensure thread safety during the clone operation
// and sends progress messages through the provided messager.
func (g *gitClient) runClone(ctx context.Context, o *CloneOptions, messager eventer.Messager) {
	// close the messager channel when the function exits
	defer messager.Close()

	if o == nil {
		logger.Error("clone options are required")
		return
	}

	// Lock the mutex to ensure thread safety during the clone operation
	// it is closed after the clone operation is complete
	// to allow other operations to proceed
	// while the clone is in progress
	g.mutex.Lock()
	defer g.mutex.Unlock()

	repositoryPath := filepath.Join(g.basePath, o.Path)

	// Error if repositoryPath is existing
	if _, err := os.Stat(repositoryPath); err == nil {
		messager.Send("error", "repository already exists")
		return
	}
	logger.Infof("cloning repository from %s to %s", o.RepoUrl, repositoryPath)

	repo, err := git.PlainCloneContext(ctx, repositoryPath, false, &git.CloneOptions{
		URL:      o.RepoUrl,
		Progress: messager,
		Auth: &http.BasicAuth{
			Username: o.User,
			Password: o.Token,
		},
	})
	if err != nil {
		messager.Send(eventer.StatusError, err.Error())
		return
	}

	// Get the head branch
	head, err := getHead(repo)
	if err != nil {
		messager.Send(eventer.StatusError, err.Error())
		return
	}
	messager.Send(eventer.StatusSuccess, toBranch(head, true))

	logger.Infof("cloned repository successfully into '%s' with head '%s'", repositoryPath, head.String())
}
