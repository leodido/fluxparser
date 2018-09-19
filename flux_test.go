package fluxparser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/ast/asttest"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    *ast.Program
		wantErr bool
	}{
		// 		{
		// 			name: "from with join with complex expression",
		// 			raw: `
		// a = from(db:"Flux")
		// 	|> filter(fn: (r) => r["_measurement"] == "a")
		// 	|> range(start:-1h)

		// b = from(db:"Flux")
		// 	|> filter(fn: (r) => r["_measurement"] == "b")
		// 	|> range(start:-1h)

		// join(tables:[a,b], on:["t1"], fn: (a,b) => (a["_field"] - b["_field"]) / b["_field"])
		// `,
		// 		},
		// {
		// 	name: "empty block statement",
		// 	raw:  `{}`,
		// },
		// {
		// 	name: "one block statement",
		// 	raw:  `{ }`,
		// },
		// {
		// 	name: "nest block statements",
		// 	raw:  `{{}}`,
		// },
		// {
		// 	name: "use variable to declare something",
		// 	raw:  `howdy = "1"`,
		// 	want: &ast.Program{
		// 		Body: []ast.Statement{
		// 			&ast.VariableDeclaration{
		// 				Declarations: []*ast.VariableDeclarator{{
		// 					ID:   &ast.Identifier{Name: "howdy"},
		// 					Init: &ast.StringLiteral{Value: "1"},
		// 				}},
		// 			},
		// 		},
		// 	},
		// },

		// {
		// 	name: "variable declaration",
		// 	raw:  `leodido = `,
		// },
		// {
		// 	name: "variable declaration",
		// 	raw:  `{docmerlin =}`,
		// },
		{
			name: "use variable to declare something",
			raw:  `from()`,
			want: &ast.Program{
				Body: []ast.Statement{
					&ast.ExpressionStatement{
						Expression: &ast.CallExpression{
							Callee: &ast.Identifier{
								Name: "from",
							},
						},
					},
				},
			},
		},
	}

	mach := NewMachine()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()

			if !tt.wantErr {
				got := mach.Parse([]byte(tt.raw))
				if !cmp.Equal(tt.want, got, asttest.CompareOptions...) {
					t.Errorf("parser.NewAST() = -want/+got %s", cmp.Diff(tt.want, got, asttest.CompareOptions...))
				}
			}

		})
	}
}

func BenchmarkParse(b *testing.B) {
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
		mach := NewMachine()
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				mach.Parse([]byte(tt.raw))
			}
		})
	}
}
