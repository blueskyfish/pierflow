package gitter

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"pierflow/internal/logger"
	"strings"
)

func getHead(repo *git.Repository) (plumbing.ReferenceName, error) {
	head, err := repo.Head()
	if err != nil {
		return "", err
	}
	return head.Name(), nil
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

func adjustBranchName(branch string) string {
	if branch == "" {
		return "main"
	}
	index := strings.Index(branch, "/")
	if index >= 0 {
		branch = branch[index+1:]
	}
	return branch
}
