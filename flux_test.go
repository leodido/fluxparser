package fluxparser

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
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
		{
			name: "empty block statement",
			raw:  `{}`,
		},
		{
			name: "one block statement",
			raw:  `{ }`,
		},
		{
			name: "nest block statements",
			raw:  `{{}}`,
		},
	}

	mach := NewMachine()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()

			fmt.Println(mach.Parse([]byte(tt.raw)))

		})
	}
}
