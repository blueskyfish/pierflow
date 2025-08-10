package gitter

import (
	"pierflow/internal/helper"

	"github.com/go-git/go-git/v5/plumbing"
)

func toBranch(b plumbing.ReferenceName, active bool) *Branch {
	return &Branch{
		Branch: b.Short(),
		Place:  helper.IIF(b.IsBranch(), BranchPlaceLocal, BranchPlaceRemote),
		Active: active,
	}
}
