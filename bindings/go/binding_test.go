package tree_sitter_imhex_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_imhex "github.com/tree-sitter/tree-sitter-imhex/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_imhex.Language())
	if language == nil {
		t.Errorf("Error loading Imhex grammar")
	}
}
