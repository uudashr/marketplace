package store_test

import (
	"testing"

	"github.com/uudashr/marketplace/internal/store"
)

func TestStore(t *testing.T) {
	cases := map[string]struct {
		id        string
		name      string
		expectErr bool
	}{
		"Default": {
			id:   "an-id",
			name: "SuperMart",
		},

		"Empty ID": {
			id:        "",
			name:      "SuperMart",
			expectErr: true,
		},

		"Empty Name": {
			id:        "an-id",
			name:      "",
			expectErr: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			s, err := store.New(c.id, c.name)
			if c.expectErr {
				if err == nil {
					t.Fatal("Expect err")
				}
				return
			}

			if err != nil {
				t.Fatal("err:", err)
			}

			if got, want := s.ID(), c.id; got != want {
				t.Errorf("ID got: %q, want: %q", got, want)
			}

			if got, want := s.Name(), c.name; got != want {
				t.Errorf("Name got: %q, want: %q", got, want)
			}
		})
	}
}
