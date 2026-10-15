package greet

import "testing"

func TestGreeting(t *testing.T) {
	cases := []struct {
		name, in, want string
		wantErr        error
	}{
		{"plain", "Ada", "Hello, Ada!", nil},
		{"trimmed", "  Ada  ", "Hello, Ada!", nil},
		{"empty", "", "", ErrEmptyName},
		{"blank", "   ", "", ErrEmptyName},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Greeting(c.in)
			if err != c.wantErr || got != c.want {
				t.Fatalf("Greeting(%q) = %q, %v; want %q, %v", c.in, got, err, c.want, c.wantErr)
			}
		})
	}
}
