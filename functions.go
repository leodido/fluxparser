package fluxparser

import (
	"strconv"

	"github.com/influxdata/flux/ast"
)

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

func getDuration(bytes []byte, unit string) ast.Duration {
	v1 := bytes[:len(bytes)-len(unit)]
	v2, _ := strconv.Atoi(string(v1))
	if unit == "μs" {
		unit = "us"
	}
	return ast.Duration{
		Magnitude: int64(v2),
		Unit:      unit,
	}
}
