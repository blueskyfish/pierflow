package gitter

import (
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func getHead(repo *git.Repository) (plumbing.ReferenceName, error) {
	head, err := repo.Head()
	if err != nil {
		return "", err
	}
	return head.Name(), nil
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
