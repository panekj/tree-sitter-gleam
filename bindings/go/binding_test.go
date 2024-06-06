package tree_sitter_gleam_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-gleam"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_gleam.Language())
	if language == nil {
		t.Errorf("Error loading Gleam grammar")
	}
}
