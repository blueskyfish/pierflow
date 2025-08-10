package gitter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdjustBranchName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "main"},
		{"main", "main"},
		{"origin/feature/branch", "feature/branch"},
		{"origin/bugfix/issue-123", "bugfix/issue-123"},
	}

	for _, test := range tests {
		result := adjustBranchName(test.input)
		assert.Equal(t, test.expected, result)
	}
}
