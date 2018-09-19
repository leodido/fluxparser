package fluxparser

import "github.com/influxdata/flux/ast"

func base(text []byte, line, col int) *ast.BaseNode {
	str := string(text)

	return &ast.BaseNode{
		Loc: &ast.SourceLocation{
			Start: ast.Position{
				Line:   line,
				Column: col - len(text),
			},
			End: ast.Position{
				Line:   line,
				Column: col,
			},
			Source: &str,
		},
	}
}
