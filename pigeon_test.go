package fluxparser

import (
	"testing"

	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/parser"
)

func BenchmarkPigeonParse(b *testing.B) {
	tests := []struct {
		name    string
		raw     string
		want    *ast.Program
		wantErr bool
	}{
		{
			name: "use variable to declare something",
			raw:  `howdy = "1"`,
			want: &ast.Program{
				Body: []ast.Statement{
					&ast.VariableDeclaration{
						Declarations: []*ast.VariableDeclarator{{
							ID:   &ast.Identifier{Name: "howdy"},
							Init: &ast.StringLiteral{Value: "1"},
						}},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				parser.NewAST(tt.raw)
			}
		})
	}
}
