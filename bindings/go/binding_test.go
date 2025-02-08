package tree_sitter_groovy_test

import (
	"testing"

	tree_sitter_groovy "github.com/herbygillot/tree-sitter-groovy/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_groovy.Language())
	if language == nil {
		t.Errorf("Error loading JavaScript grammar")
	}
}
