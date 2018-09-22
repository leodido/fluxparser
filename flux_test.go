package fluxparser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/ast/asttest"
	"github.com/influxdata/flux/parser"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    *ast.Program
		wantErr bool
	}{
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
		// 	name: "variable declaration",
		// 	raw:  `leodido = `,
		// },
		// {
		// 	name: "variable declaration",
		// 	raw:  `{docmerlin = "emrys"}`,
		// },

		{
			name: "use variable to declare something",
			raw:  `howdy = "literal"`,
			want: &ast.Program{
				Body: []ast.Statement{
					&ast.VariableDeclaration{
						Declarations: []*ast.VariableDeclarator{{
							ID:   &ast.Identifier{Name: "howdy"},
							Init: &ast.StringLiteral{Value: "literal"},
						}},
					},
				},
			},
		},
		{
			name: "use variable to declare something, without spaces",
			raw:  `howdy="1"`,
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
		// {
		// 	name: "from function",
		// 	raw:  `from()`,
		// 	want: &ast.Program{
		// 		Body: []ast.Statement{
		// 			&ast.ExpressionStatement{
		// 				Expression: &ast.CallExpression{
		// 					Callee: &ast.Identifier{
		// 						Name: "from",
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		{
			name: "duration literal, all units",
			raw:  `dur = 1y3mo2w1d4h1m30s5ms2μs70ns`,
			want: &ast.Program{
				Body: []ast.Statement{&ast.VariableDeclaration{
					Declarations: []*ast.VariableDeclarator{{
						ID: &ast.Identifier{Name: "dur"},
						Init: &ast.DurationLiteral{
							Values: []ast.Duration{
								{Magnitude: 1, Unit: "y"},
								{Magnitude: 3, Unit: "mo"},
								{Magnitude: 2, Unit: "w"},
								{Magnitude: 1, Unit: "d"},
								{Magnitude: 4, Unit: "h"},
								{Magnitude: 1, Unit: "m"},
								{Magnitude: 30, Unit: "s"},
								{Magnitude: 5, Unit: "ms"},
								{Magnitude: 2, Unit: "us"},
								{Magnitude: 70, Unit: "ns"},
							},
						},
					}},
				}},
			},
		},
		{
			name: "duration literal, months",
			raw:  `dur = 6mo`,
			want: &ast.Program{
				Body: []ast.Statement{&ast.VariableDeclaration{
					Declarations: []*ast.VariableDeclarator{{
						ID: &ast.Identifier{Name: "dur"},
						Init: &ast.DurationLiteral{
							Values: []ast.Duration{
								{Magnitude: 6, Unit: "mo"},
							},
						},
					}},
				}},
			},
		},
		{
			name: "duration literal, milliseconds",
			raw:  `dur = 500ms`,
			want: &ast.Program{
				Body: []ast.Statement{&ast.VariableDeclaration{
					Declarations: []*ast.VariableDeclarator{{
						ID: &ast.Identifier{Name: "dur"},
						Init: &ast.DurationLiteral{
							Values: []ast.Duration{
								{Magnitude: 500, Unit: "ms"},
							},
						},
					}},
				}},
			},
		},
		{
			name: "duration literal, months, minutes, milliseconds",
			raw:  `dur = 6mo30m500ms`,
			want: &ast.Program{
				Body: []ast.Statement{&ast.VariableDeclaration{
					Declarations: []*ast.VariableDeclarator{{
						ID: &ast.Identifier{Name: "dur"},
						Init: &ast.DurationLiteral{
							Values: []ast.Duration{
								{Magnitude: 6, Unit: "mo"},
								{Magnitude: 30, Unit: "m"},
								{Magnitude: 500, Unit: "ms"},
							},
						},
					}},
				}},
			},
		},
		{
			name:    "duration literal, units unordered are not allowed",
			raw:     `dur = 6mo30m3mo`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mach := NewMachine()
			if !tt.wantErr {
				got := mach.Parse([]byte(tt.raw))
				if !cmp.Equal(tt.want, got, asttest.CompareOptions...) {
					t.Errorf("-want/+got %s", cmp.Diff(tt.want, got, asttest.CompareOptions...))
				}
			}

		})
	}
}

var btests = []struct {
	name string
	raw  string
}{
	{
		name: "use variable to declare something",
		raw:  `howdy = "literal"`,
	},
	{
		name: "use variable to declare something without spaces",
		raw:  `howdy="1"`,
	},
	{
		name: "from function",
		raw:  `from()`,
	},
	{
		name: "duration literal, all units",
		raw:  `dur = 1y3mo2w1d4h1m30s5ms2μs70ns`,
	},
}

func BenchmarkParse(b *testing.B) {
	for _, tt := range btests {
		mach := NewMachine()
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				mach.Parse([]byte(tt.raw))
			}
		})
	}
}

func BenchmarkPigeonParse(b *testing.B) {
	for _, tt := range btests {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				parser.NewAST(tt.raw)
			}
		})
	}
}
