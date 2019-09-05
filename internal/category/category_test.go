package category_test

import (
	"testing"

	"github.com/uudashr/marketplace/internal/category"
)

func TestCategory(t *testing.T) {
	cases := map[string]struct {
		ID          string
		Name        string
		ExpectError bool
	}{
		"Default": {
			ID:   category.NextID(),
			Name: "Utilities",
		},
		"Empty ID": {
			ID:          "",
			Name:        "Utilities",
			ExpectError: true,
		},
		"Empty Name": {
			ID:          category.NextID(),
			Name:        "",
			ExpectError: true,
		},
	}

	for k, c := range cases {
		t.Run(k, func(t *testing.T) {
			cat, err := category.New(c.ID, c.Name)
			if c.ExpectError {
				if err == nil {
					t.Errorf("Expecting error for case: %q", k)
				}
				return
			}

			if err != nil {
				t.Fatalf("case: %q, err: %v", k, err)
			}

			if got, want := cat.ID(), c.ID; got != want {
				t.Errorf("ID got: %q, want: %q, case: %q", got, want, k)
			}

			if got, want := cat.Name(), c.Name; got != want {
				t.Errorf("Name got: %q, want: %q, case: %q", got, want, k)
			}
		})
	}
}
