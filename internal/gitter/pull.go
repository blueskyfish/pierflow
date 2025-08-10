package gitter

import (
	"context"
	"errors"
	"path/filepath"
	"pierflow/internal/logger"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func (g *gitClient) Pull(ctx context.Context, path string, o *PullOptions) (string, *Branch, error) {
	if o == nil {
		return "", nil, errors.New("pull options are required")
	}

	reposPath := filepath.Join(g.basePath, path)
	logger.Infof("pull repository '%s'", reposPath)

	// Open git repository
	repo, err := git.PlainOpen(reposPath)
	if err != nil {
		return "", nil, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return "", nil, err
	}

	progressor := newProgressor()

	head, err := getHead(repo)
	if err != nil {
		return "", nil, err
	}

	// Pull the latest changes
	err = worktree.PullContext(ctx, &git.PullOptions{
		Auth:          &http.BasicAuth{Username: o.User, Password: o.Token},
		RemoteName:    git.DefaultRemoteName,
		ReferenceName: head,
		RemoteURL:     o.GitUrl,
		Progress:      progressor,
		Force:         true,
		SingleBranch:  true,
	})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return "", nil, err
	}

	return progressor.String(), toBranch(head, true), nil
}
