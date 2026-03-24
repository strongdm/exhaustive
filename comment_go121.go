//go:build go1.21

package sdmexhaustive

import (
	"go/ast"
)

func isGeneratedFile(file *ast.File) bool {
	return ast.IsGenerated(file)
}
