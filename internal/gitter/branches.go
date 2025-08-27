package gitter

import (
	"context"
	"errors"
	"path/filepath"
	"pierflow/internal/eventer"
	"pierflow/internal/logger"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type BranchPlace int

const (
	BranchPlaceUnknown BranchPlace = -1
	BranchPlaceLocal   BranchPlace = 0
	BranchPlaceRemote  BranchPlace = 1

	BranchPlaceLocalName   = "local"
	BranchPlaceRemoteName  = "remote"
	BranchPlaceUnknownName = "unknown"
)

type BranchOptions struct {
	Refresh bool
	User    string
	Token   string
	Path    string
	Prune   bool
}

func (p BranchPlace) String() string {
	switch p {
	case BranchPlaceLocal:
		return BranchPlaceLocalName
	case BranchPlaceRemote:
		return BranchPlaceRemoteName
	default:
		return BranchPlaceUnknownName
	}
}

func StringBranchPlace(s string) (BranchPlace, error) {
	switch s {
	case BranchPlaceLocalName:
		return BranchPlaceLocal, nil
	case BranchPlaceRemoteName:
		return BranchPlaceRemote, nil
	default:
		return BranchPlaceUnknown, errors.New("unknown branch place")
	}
}

type Branch struct {
	Branch string
	Place  BranchPlace
	Path   string
	Active bool
}

func (g *gitClient) BranchList(ctx context.Context, options *BranchOptions, messager eventer.Messager) {
	go g.runBranchList(ctx, options, messager)
}

func (g *gitClient) runBranchList(ctx context.Context, o *BranchOptions, messager eventer.Messager) {
	// close the messager channel when the function exits
	defer messager.Closing()

	g.mutex.Lock()
	defer g.mutex.Unlock()

	if o == nil {
		logger.Error("branch options are required")
		return
	}

	reposPath := filepath.Join(g.basePath, o.Path)

	// Open git repository
	repo, err := git.PlainOpen(reposPath)
	if err != nil {
		_ = messager.Send(eventer.StatusError, err.Error())
		return
	}

	if o.Refresh {
		err = repo.FetchContext(ctx, &git.FetchOptions{
			Force: true,
			Auth: &http.BasicAuth{
				Username: o.User,
				Password: o.Token,
			},
			Prune:    o.Prune,
			Progress: messager,
		})
		if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			_ = messager.Send(eventer.StatusError, err.Error())
			return
		}
		logger.Infof("fetch branches in repository '%s'", reposPath)
	}

	// Get all references in the repository
	refs, err := repo.References()
	if err != nil {
		_ = messager.Send(eventer.StatusError, "Repository without branches")
		return
	}

	// Get the current branch
	head, err := getHead(repo)
	if err != nil {
		_ = messager.Send(eventer.StatusError, "Repository without header")
		return
	}
	currentBranch := head.Short()
	logger.Infof("branch list with current branch: %s", currentBranch)

	var list []Branch
	_ = refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsBranch() {
			list = append(list, Branch{
				Branch: ref.Name().Short(),
				Place:  BranchPlaceLocal,
				Active: ref.Name().Short() == currentBranch,
			})
		} else if ref.Name().IsRemote() {
			list = append(list, Branch{
				Branch: ref.Name().Short(),
				Place:  BranchPlaceRemote,
				Active: ref.Name().Short() == currentBranch,
			})
		}
		return nil
	})

	if len(list) == 0 {
		_ = messager.Send(eventer.StatusInfo, "Repository empty branches")
		return
	}

	_ = messager.Send(eventer.StatusSuccess, list)
}
