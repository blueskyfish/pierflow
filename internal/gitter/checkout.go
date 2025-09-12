package gitter

import (
	"fmt"
	"path/filepath"

	"github.com/blueskyfish/pierflow/internal/eventer"
	"github.com/blueskyfish/pierflow/internal/logger"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type CheckoutOptions struct {
	Branch string
	Place  string
	Path   string
}

func (g *gitClient) Checkout(o *CheckoutOptions, messager eventer.Messager) {
	go g.runCheckout(o, messager)
}

func (g *gitClient) runCheckout(o *CheckoutOptions, messager eventer.Messager) {
	defer messager.Close()

	if o == nil {
		logger.Error("checkout options are required")
		return
	}

	place, err := StringBranchPlace(o.Place)
	if err != nil {
		messager.Send(eventer.StatusError, err.Error())
		return
	}

	branch := o.Branch
	if place == BranchPlaceRemote {
		branch = adjustBranchName(branch)
	}

	reposPath := filepath.Join(g.basePath, o.Path)
	logger.Infof("checkout branch '%s' in repository '%s'", o.Branch, reposPath)

	// Open git repository
	repo, err := git.PlainOpen(reposPath)
	if err != nil {
		messager.Send(eventer.StatusError, err.Error())
		return
	}

	// Chec kout the specified branch
	worktree, err := repo.Worktree()
	if err != nil {
		messager.Send(eventer.StatusError, err.Error())
		return
	}

	// Get the current branch
	head, err := getHead(repo)
	if err != nil {
		messager.Send(eventer.StatusError, err.Error())
		return
	}

	branchRef := plumbing.NewBranchReferenceName(branch)
	ref, err := repo.Reference(branchRef, true)
	branchExist := err == nil || (ref != nil && ref.Name().Short() == head.Short())

	messager.Send(eventer.StatusInfo, fmt.Sprintf("start checkout branch '%s'", branch))
	err = checkoutBranch(worktree, branchExist, branchRef)
	if err != nil {
		messager.Send(eventer.StatusError, err.Error())
		return
	}
	messager.Send(eventer.StatusInfo, fmt.Sprintf("checkout branch '%s' success", branch))

	head, err = getHead(repo)
	if err != nil {
		messager.Send(eventer.StatusError, err.Error())
		return
	}
	messager.Send(eventer.StatusSuccess, toBranch(head, true))
}

func checkoutBranch(worktree *git.Worktree, exist bool, branchRef plumbing.ReferenceName) error {
	if exist {
		logger.Infof("checkout branch '%s'...", branchRef.Short())
		err := worktree.Checkout(&git.CheckoutOptions{
			Branch: branchRef,
			Create: false,
			Force:  true,
			Keep:   true,
		})
		if err != nil {
			return err
		}
	} else {
		// Create and checkout a new branch
		logger.Infof("create and checkout new branch '%s'...", branchRef.Short())
		err := worktree.Checkout(&git.CheckoutOptions{
			Branch: branchRef,
			Create: true,
			Force:  true,
			Keep:   true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
