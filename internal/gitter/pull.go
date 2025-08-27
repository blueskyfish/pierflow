package gitter

import (
	"context"
	"errors"
	"path/filepath"
	"pierflow/internal/eventer"
	"pierflow/internal/logger"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type PullOptions struct {
	User   string
	Token  string
	GitUrl string
	Path   string
}

func (g *gitClient) Pull(ctx context.Context, o *PullOptions, messager eventer.Messager) {
	go g.runPull(ctx, o, messager)
}

func (g *gitClient) runPull(ctx context.Context, o *PullOptions, messager eventer.Messager) {
	// close the messager channel when the function exits
	defer messager.Closing()

	if o == nil {
		logger.Error("pull options are required")
		return
	}

	// Lock the mutex to ensure thread safety during the pull operation
	// it is closed after the pull operation is complete
	// to allow other operations to proceed
	g.mutex.Lock()
	defer g.mutex.Unlock()

	repositoryPath := filepath.Join(g.basePath, o.Path)

	// Open git repository
	repo, err := git.PlainOpen(repositoryPath)
	if err != nil {
		_ = messager.Send(eventer.StatusError, err.Error())
		return
	}

	worktree, err := repo.Worktree()
	if err != nil {
		_ = messager.Send(eventer.StatusError, err.Error())
		return
	}

	head, err := getHead(repo)
	if err != nil {
		_ = messager.Send(eventer.StatusError, err.Error())
		return
	}

	// Pull the latest changes
	err = worktree.PullContext(ctx, &git.PullOptions{
		Auth:          &http.BasicAuth{Username: o.User, Password: o.Token},
		RemoteName:    git.DefaultRemoteName,
		ReferenceName: head,
		RemoteURL:     o.GitUrl,
		Progress:      messager,
		Force:         true,
		SingleBranch:  true,
	})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		_ = messager.Send(eventer.StatusError, err.Error())
		return
	}

	head, err = getHead(repo)
	if err != nil {
		_ = messager.Send(eventer.StatusError, err.Error())
		return
	}

	_ = messager.Send(eventer.StatusSuccess, toBranch(head, true))
}
