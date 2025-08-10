package gitter

import (
	"context"
	"errors"
	"path/filepath"
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
	Branch string      `json:"branch"`
	Place  BranchPlace `json:"place"`
	Active bool        `json:"active"`
}

func (g *gitClient) BranchList(ctx context.Context, path string, options *BranchOptions) (string, []Branch, error) {

	if options == nil {
		options = &BranchOptions{Refresh: false, User: "", Token: "", Prune: false}
	}

	reposPath := filepath.Join(g.basePath, path)

	// Open git repository
	repo, err := git.PlainOpen(reposPath)
	if err != nil {
		return "", nil, err
	}

	progressor := newProgressor()

	if options.Refresh {
		err = repo.FetchContext(ctx, &git.FetchOptions{
			Force: true,
			Auth: &http.BasicAuth{
				Username: options.User,
				Password: options.Token,
			},
			Prune:    options.Prune,
			Progress: progressor,
		})
		if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			return "Refresh is failed", nil, err
		}
		logger.Infof("fetch branches in repository '%s'", reposPath)
	}

	// Get all references in the repository
	refs, err := repo.References()
	if err != nil {
		return "Repository without branches", nil, err
	}

	// Get the current branch
	head, err := getHead(repo)
	if err != nil {
		return "Repository without header", nil, err
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
		return "Repository empty branches", []Branch{}, nil
	}
	return progressor.String(), list, nil
}

func (g *gitClient) CheckoutBranch(path string, o *CheckoutOptions) (*Branch, error) {
	if o == nil {
		return nil, errors.New("checkout options are required")
	}

	place, err := StringBranchPlace(o.Place)
	if err != nil {
		return nil, err
	}

	branch := o.Branch
	if place == BranchPlaceRemote {
		branch = adjustBranchName(branch)
	}

	reposPath := filepath.Join(g.basePath, path)
	logger.Infof("checkout branch '%s' in repository '%s'", o.Branch, reposPath)

	// Open git repository
	repo, err := git.PlainOpen(reposPath)
	if err != nil {
		return nil, err
	}

	// Checkout the specified branch
	worktree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	// Get the current branch
	head, err := getHead(repo)
	if err != nil {
		return nil, err
	}

	branchRef := plumbing.NewBranchReferenceName(branch)
	ref, err := repo.Reference(branchRef, true)
	branchExist := err == nil || (ref != nil && ref.Name().Short() == head.Short())

	err = checkoutBranch(worktree, branchExist, branchRef)
	if err != nil {
		return nil, err
	}

	head, err = getHead(repo)
	if err != nil {
		return nil, err
	}

	return toBranch(head, true), nil
}
