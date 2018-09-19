package fluxparser

import (
	"testing"

	"github.com/influxdata/flux/parser"
)

func BenchmarkPigeonParse(b *testing.B) {
	tests := []struct {
		name string
		raw  string
	}{
		{
			name: "use variable to declare something",
			raw:  `howdy = "1"`,
		},
		{
			name: "use variable to declare something without spaces",
			raw:  `howdy="1"`,
		},
		{
			name: "from function",
			raw:  `from()`,
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
