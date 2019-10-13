package testutils

import (
	"github.com/elliotcourant/melogale/pkg/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func MustParse(t *testing.T, query string) ast.SyntaxTree {
	tree, err := ast.Parse(query)
	if !assert.NoError(t, err, "failed to parse query") {
		panic(err)
	}
	return tree
}
